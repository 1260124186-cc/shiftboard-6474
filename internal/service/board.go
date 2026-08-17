package service

import (
	"context"
	"fmt"

	"shiftboard/internal/model"
	"shiftboard/internal/plan"
	"shiftboard/internal/report"
	"shiftboard/internal/store"
)

type Board struct {
	repository *store.Repository
	policies   []model.ZonePolicy
}

func NewBoard() *Board {
	return NewBoardWithPolicies([]model.ZonePolicy{
		{Zone: "north", MaxPerShift: 2},
		{Zone: "south", MaxPerShift: 2},
	})
}

func NewBoardWithPolicies(policies []model.ZonePolicy) *Board {
	return &Board{
		repository: store.NewRepository(),
		policies:   clonePolicies(policies),
	}
}

func clonePolicies(policies []model.ZonePolicy) []model.ZonePolicy {
	cloned := make([]model.ZonePolicy, len(policies))
	for i := range policies {
		cloned[i] = policies[i].Clone()
	}
	return cloned
}

func (board *Board) Import(ctx context.Context, orders []model.WorkOrder) error {
	for _, order := range orders {
		if err := ctx.Err(); err != nil {
			return err
		}
		if err := board.repository.Save(order); err != nil {
			return fmt.Errorf("import work order %s: %w", order.ID, err)
		}
	}
	return nil
}

func (board *Board) Assign() error {
	orders := board.repository.List()
	assignments, err := plan.BuildAssignments(orders, board.policies)
	if err != nil {
		return err
	}
	for _, assignment := range assignments {
		order, err := board.repository.Get(assignment.WorkOrderID)
		if err != nil {
			return err
		}
		order.Stage = model.StageAssigned
		order.Metadata = map[string]string{"shift": assignment.Shift}
		if err := board.repository.Replace(order); err != nil {
			return err
		}
	}
	return nil
}

func (board *Board) Report() string {
	orders := board.repository.List()
	assignments, err := plan.BuildAssignments(orders, board.policies)
	if err != nil {
		return "work orders: 0\n"
	}
	return report.Format(orders, assignments)
}

func (board *Board) Orders() []model.WorkOrder {
	return board.repository.List()
}
