package report

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"shiftboard/internal/model"
	"shiftboard/internal/plan"
)

func FormatImportError(err error) string {
	if errors.Is(err, model.ErrImportCanceled) {
		return "import canceled without changing the board\n"
	}
	return "import failed\n"
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
