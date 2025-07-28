package cost

import "strings"

// Estimator provides cost estimates for resources.
type Estimator interface {
	EstimateCost(resourceType string, count int) float64
}

type providerEstimator struct {
	costs map[string]float64
}

func (e providerEstimator) EstimateCost(resourceType string, count int) float64 {
	key := strings.ToLower(resourceType)
	if c, ok := e.costs[key]; ok {
		return c * float64(count)
	}
	return 0.0
}

func NewEstimator(provider string) Estimator {
	switch provider {
	case "aws":
		return providerEstimator{costs: awsCosts}
	case "gcp":
		return providerEstimator{costs: gcpCosts}
	case "azure":
		return providerEstimator{costs: azureCosts}
	default:
		return providerEstimator{costs: map[string]float64{}}
	}
}
