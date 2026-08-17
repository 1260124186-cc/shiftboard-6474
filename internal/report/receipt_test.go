package report

import "testing"

func TestDispatchReceiptIncludesReadyStatus(t *testing.T) {
	const want = "dispatched work orders: 2\nstatus: ready\n"
	if got := FormatDispatchReceipt(2); got != want {
		t.Fatalf("receipt = %q, want %q", got, want)
	}
}
