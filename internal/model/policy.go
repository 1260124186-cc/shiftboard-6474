package model

import "slices"

type ZonePolicy struct {
	Zone          string
	MaxPerShift   int
	RequiredTags  []string
	AllowOverflow bool
}

func (policy ZonePolicy) Clone() ZonePolicy {
	clone := policy
	clone.RequiredTags = slices.Clone(policy.RequiredTags)
	return clone
}
