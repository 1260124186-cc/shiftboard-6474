package report

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"shiftboard/internal/model"
	"shiftboard/internal/plan"
)

func FormatPlanningError(err error) string {
	if errors.Is(err, model.ErrMissingZonePolicy) {
		return "planning blocked: missing zone policy\n"
	}
	return "planning unavailable\n"
}

func Format(orders []model.WorkOrder, assignments []plan.Assignment) string {
	counts := make(map[string]int)
	for _, assignment := range assignments {
		counts[assignment.Shift]++
	}

	shifts := make([]string, 0, len(counts))
	for shift := range counts {
		shifts = append(shifts, shift)
	}
	sort.Strings(shifts)

	lines := []string{fmt.Sprintf("work orders: %d", len(orders))}
	for _, shift := range shifts {
		lines = append(lines, fmt.Sprintf("%s: %d", shift, counts[shift]))
	}
	return strings.Join(lines, "\n") + "\n"
}
