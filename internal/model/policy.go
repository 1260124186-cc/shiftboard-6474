package model

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
