package model

import (
	"fmt"
	"maps"
	"slices"
)

type Stage string

const (
	StageQueued   Stage = "queued"
	StageAssigned Stage = "assigned"
	StageComplete Stage = "complete"
)

type WorkOrder struct {
	ID       string
	Title    string
	Zone     string
	Stage    Stage
	Tags     []string
	Notes    []string
	Metadata map[string]string
}

func (order WorkOrder) Validate() error {
	if order.ID == "" {
		return fmt.Errorf("work order id is required")
	}
	if order.Title == "" {
		return fmt.Errorf("work order %s title is required", order.ID)
	}
	if order.Zone == "" {
		return fmt.Errorf("work order %s zone is required", order.ID)
	}
	return nil
}

func (order WorkOrder) Clone() WorkOrder {
	clone := order
	clone.Tags = slices.Clone(order.Tags)
	clone.Notes = slices.Clone(order.Notes)
	if order.Metadata != nil {
		clone.Metadata = maps.Clone(order.Metadata)
	}
	return clone
}
