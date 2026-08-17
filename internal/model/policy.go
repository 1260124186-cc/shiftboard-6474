package model

import (
	"errors"
	"fmt"
)

var ErrMissingZonePolicy = errors.New("missing zone policy")

// MissingZonePolicyError 表示某个区域缺少排班规则，归类到 ErrMissingZonePolicy，
// 以便跨排班与报表流程通过 errors.Is 可靠识别。
type MissingZonePolicyError struct {
	Zone string
}

func (err MissingZonePolicyError) Error() string {
	return fmt.Sprintf("missing policy for zone %s", err.Zone)
}

// Is 让 errors.Is(err, ErrMissingZonePolicy) 对该业务错误返回 true。
func (err MissingZonePolicyError) Is(target error) bool {
	return target == ErrMissingZonePolicy
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
