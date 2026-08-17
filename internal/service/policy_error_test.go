package service_test

import (
	"context"
	"errors"
	"testing"

	"shiftboard/internal/model"
	"shiftboard/internal/service"
)

func TestMissingZonePolicyKeepsItsClassification(t *testing.T) {
	board := service.NewBoardWithPolicies([]model.ZonePolicy{{Zone: "north", MaxPerShift: 1}})
	if err := board.Import(context.Background(), []model.WorkOrder{{
		ID: "WO-uncovered", Title: "Inspect south dock", Zone: "south",
	}}); err != nil {
		t.Fatalf("import: %v", err)
	}

	err := board.Assign()
	if !errors.Is(err, model.ErrMissingZonePolicy) {
		t.Fatalf("assignment error did not preserve the policy classification: %v", err)
	}
	if got := board.Report(); got != "planning blocked: missing zone policy\n" {
		t.Fatalf("report = %q", got)
	}
}
