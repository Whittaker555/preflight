package cost

import (
	"strings"
)

func EstimateCost(resourceType string, count int) float64 {
	// Normalise (Terraform sometimes returns type with provider prefix)
	key := strings.ToLower(resourceType)
	if cost, ok := ResourceCosts[key]; ok {
		return cost * float64(count)
	}
	return 0.0
}
