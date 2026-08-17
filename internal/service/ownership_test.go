package service_test

import (
	"context"
	"testing"

	"shiftboard/internal/model"
	"shiftboard/internal/service"
)

func TestBoardOwnsImportedMutableData(t *testing.T) {
	policies := []model.ZonePolicy{{
		Zone:         "north",
		MaxPerShift:  1,
		RequiredTags: []string{"safety"},
	}}
	order := model.WorkOrder{
		ID:       "WO-ownership",
		Title:    "Inspect dock",
		Zone:     "north",
		Tags:     []string{"safety"},
		Metadata: map[string]string{"source": "tablet"},
	}
	board := service.NewBoardWithPolicies(policies)
	if err := board.Import(context.Background(), []model.WorkOrder{order}); err != nil {
		t.Fatalf("import: %v", err)
	}

	order.Tags[0] = "inventory"
	order.Metadata["source"] = "rewritten"
	policies[0].RequiredTags[0] = "emergency"

	if err := board.Assign(); err != nil {
		t.Fatalf("assign after caller mutations: %v", err)
	}
	got := board.Orders()[0]
	if got.Tags[0] != "safety" {
		t.Fatalf("board retained caller-owned mutable values: %#v", got)
	}
}
