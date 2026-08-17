package model

type ZonePolicy struct {
	Zone          string
	MaxPerShift   int
	RequiredTags  []string
	AllowOverflow bool
}

func (policy ZonePolicy) Clone() ZonePolicy {
	clone := policy
	clone.RequiredTags = policy.RequiredTags
	return clone
}
