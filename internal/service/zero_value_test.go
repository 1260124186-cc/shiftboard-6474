package service_test

import (
	"context"
	"testing"

	"shiftboard/internal/model"
	"shiftboard/internal/service"
)

func TestZeroValueBoardCanImportAndListWorkOrders(t *testing.T) {
	var board service.Board
	defer func() {
		if recovered := recover(); recovered != nil {
			t.Fatalf("zero-value board panicked: %v", recovered)
		}
	}()

	if err := board.Import(context.Background(), []model.WorkOrder{{
		ID: "WO-zero", Title: "Inspect dock", Zone: "north",
	}}); err != nil {
		t.Fatalf("import: %v", err)
	}
	if got := len(board.Orders()); got != 1 {
		t.Fatalf("orders = %d, want 1", got)
	}
}
