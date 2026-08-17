package plan

import (
	"fmt"
	"sort"

	"shiftboard/internal/model"
)

type Assignment struct {
	WorkOrderID string
	Shift       string
	Zone        string
}

func BuildAssignments(orders []model.WorkOrder, policies []model.ZonePolicy) ([]Assignment, error) {
	limits := make(map[string]int, len(policies))
	for _, policy := range policies {
		if policy.Zone == "" || policy.MaxPerShift < 1 {
			return nil, fmt.Errorf("invalid zone policy")
		}
		limits[policy.Zone] = policy.MaxPerShift
	}

	byZone := make(map[string]int)
	assignments := make([]Assignment, 0, len(orders))
	for _, order := range orders {
		if order.Stage != model.StageQueued {
			continue
		}
		limit, found := limits[order.Zone]
		if !found {
			return nil, fmt.Errorf("no policy for zone %s", order.Zone)
		}
		if !hasRequiredTags(order.Tags, policiesForZone(policies, order.Zone).RequiredTags) {
			return nil, fmt.Errorf("work order %s is missing a required zone tag", order.ID)
		}
		byZone[order.Zone]++
		shift := "morning"
		if byZone[order.Zone] > limit {
			shift = "afternoon"
		}
		assignments = append(assignments, Assignment{
			WorkOrderID: order.ID,
			Shift:       shift,
			Zone:        order.Zone,
		})
	}

	sort.Slice(assignments, func(left, right int) bool {
		return assignments[left].WorkOrderID < assignments[right].WorkOrderID
	})
	return assignments, nil
}

func policiesForZone(policies []model.ZonePolicy, zone string) model.ZonePolicy {
	for _, policy := range policies {
		if policy.Zone == zone {
			return policy
		}
	}
	return model.ZonePolicy{}
}

func hasRequiredTags(tags []string, required []string) bool {
	present := make(map[string]bool, len(tags))
	for _, tag := range tags {
		present[tag] = true
	}
	for _, tag := range required {
		if !present[tag] {
			return false
		}
	}
	return true
}
