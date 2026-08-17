package plan

import (
	"fmt"
	"sort"

	"shiftboard/internal/model"
)

type Assignment struct {
	WorkOrderID string
	Shift       string
	Zone        string
}

func BuildAssignments(orders []model.WorkOrder, policies []model.ZonePolicy) ([]Assignment, error) {
	limits := make(map[string]int, len(policies))
	for _, policy := range policies {
		if policy.Zone == "" || policy.MaxPerShift < 1 {
			return nil, fmt.Errorf("invalid zone policy")
		}
		limits[policy.Zone] = policy.MaxPerShift
	}

	byZone := make(map[string]int)
	assignments := make([]Assignment, 0, len(orders))
	for _, order := range orders {
		if order.Stage != model.StageQueued {
			continue
		}
		limit, found := limits[order.Zone]
		if !found {
			return nil, model.MissingZonePolicyError{Zone: order.Zone}
		}
		byZone[order.Zone]++
		shift := "morning"
		if byZone[order.Zone] > limit {
			shift = "afternoon"
		}
		assignments = append(assignments, Assignment{
			WorkOrderID: order.ID,
			Shift:       shift,
			Zone:        order.Zone,
		})
	}

	sort.Slice(assignments, func(left, right int) bool {
		return assignments[left].WorkOrderID < assignments[right].WorkOrderID
	})
	return assignments, nil
}
