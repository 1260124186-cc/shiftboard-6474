package model

import "testing"

func TestWorkOrderCloneOwnsMutableFields(t *testing.T) {
	order := WorkOrder{
		ID:       "WO-1",
		Title:    "Inspect",
		Zone:     "north",
		Tags:     []string{"daily"},
		Metadata: map[string]string{"shift": "morning"},
	}
	clone := order.Clone()
	clone.Tags[0] = "weekly"
	clone.Metadata["shift"] = "afternoon"

	if order.Tags[0] != "daily" {
		t.Fatalf("tags were shared")
	}
	if order.Metadata["shift"] != "morning" {
		t.Fatalf("metadata was shared")
	}
}
