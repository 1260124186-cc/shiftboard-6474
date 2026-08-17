package model

import (
	"errors"
	"fmt"
)

var ErrMissingZonePolicy = errors.New("missing zone policy")

type MissingZonePolicyError struct {
	Zone string
}

func (err MissingZonePolicyError) Error() string {
	return fmt.Sprintf("missing policy for zone %s", err.Zone)
}

func (err MissingZonePolicyError) Unwrap() error {
	return ErrMissingZonePolicy
}

type ZonePolicy struct {
	Zone          string
	MaxPerShift   int
	RequiredTags  []string
	AllowOverflow bool
}

func (policy ZonePolicy) Clone() ZonePolicy {
	clone := policy
	clone.RequiredTags = append([]string(nil), policy.RequiredTags...)
	return clone
}
