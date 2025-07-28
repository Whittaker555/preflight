package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whittaker555/preflight/internal/cost"
	"github.com/whittaker555/preflight/internal/models"
)

func AnalysePlan(c *gin.Context) {
	var plan models.TerraformPlan
	if err := json.NewDecoder(c.Request.Body).Decode(&plan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	c.JSON(http.StatusOK, analyse(plan))
}

// analyse generates the cost summary for a Terraform plan.
func analyse(plan models.TerraformPlan) gin.H {
	provider := cost.DetectProvider(plan)
	estimator := cost.NewEstimator(provider)

	var resources []map[string]interface{}
	var totalCost float64

	for _, rc := range plan.ResourceChanges {
		// Assume 1 resource per change for now
		costEstimate := estimator.EstimateCost(rc.Type, 1)
		totalCost += costEstimate

		resources = append(resources, map[string]interface{}{
			"address":               rc.Address,
			"type":                  rc.Type,
			"name":                  rc.Name,
			"actions":               rc.Change.Actions,
			"monthly_cost_estimate": costEstimate,
		})
	}

	return gin.H{
		"provider":                    provider,
		"total_resources":             len(resources),
		"total_monthly_cost_estimate": totalCost,
		"resources":                   resources,
	}
}
