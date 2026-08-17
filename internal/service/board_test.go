package service_test

import (
	"context"
	"testing"

	"shiftboard/internal/model"
	"shiftboard/internal/service"
)

func TestBoardImportsAssignsAndReports(t *testing.T) {
	board := service.NewBoard()
	err := board.Import(context.Background(), []model.WorkOrder{
		{ID: "WO-1", Title: "Inspect", Zone: "north"},
		{ID: "WO-2", Title: "Restock", Zone: "south"},
	})
	if err != nil {
		t.Fatalf("import: %v", err)
	}
	if err := board.Assign(); err != nil {
		t.Fatalf("assign: %v", err)
	}
	if got := board.Report(); got != "work orders: 2\n" {
		t.Fatalf("report = %q", got)
	}
	for _, order := range board.Orders() {
		if order.Stage != model.StageAssigned {
			t.Fatalf("work order %s stage = %s", order.ID, order.Stage)
		}
	}
}
