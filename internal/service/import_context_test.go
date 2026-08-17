package service_test

import (
	"context"
	"errors"
	"testing"

	"shiftboard/internal/model"
	"shiftboard/internal/report"
	"shiftboard/internal/service"
)

func TestCanceledImportDoesNotPartiallyPersistWorkOrders(t *testing.T) {
	board := service.NewBoard()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := board.Import(ctx, []model.WorkOrder{
		{ID: "WO-cancel-1", Title: "Inspect dock", Zone: "north"},
		{ID: "WO-cancel-2", Title: "Restock kits", Zone: "south"},
	})
	if !errors.Is(err, model.ErrImportCanceled) {
		t.Fatalf("import error did not preserve cancellation: %v", err)
	}
	if got := len(board.Orders()); got != 0 {
		t.Fatalf("canceled import persisted %d work orders", got)
	}
	if got := report.FormatImportError(err); got != "import canceled without changing the board\n" {
		t.Fatalf("message = %q", got)
	}
}
