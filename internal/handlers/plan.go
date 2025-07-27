package handlers

import (
	"encoding/json"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/preflight/internal/cost"
	"github.com/yourusername/preflight/internal/models"
)

func AnalysePlan(c *gin.Context) {
	var plan models.TerraformPlan
	if err := json.NewDecoder(c.Request.Body).Decode(&plan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	c.JSON(http.StatusOK, analyse(plan))
}

// UploadPlan handles multipart file uploads of Terraform plan JSON.
// The uploaded file should be provided under the "file" form field.
func UploadPlan(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "plan file required"})
		return
	}

	var opened multipart.File
	opened, err = file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to read file"})
		return
	}
	defer opened.Close()

	var plan models.TerraformPlan
	if err := json.NewDecoder(opened).Decode(&plan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid plan file"})
		return
	}

	c.JSON(http.StatusOK, analyse(plan))
}

// analyse generates the cost summary for a Terraform plan.
func analyse(plan models.TerraformPlan) gin.H {
	var resources []map[string]interface{}
	var totalCost float64

	for _, rc := range plan.ResourceChanges {
		// Assume 1 resource per change for now
		costEstimate := cost.EstimateCost(rc.Type, 1)
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
		"total_resources":             len(resources),
		"total_monthly_cost_estimate": totalCost,
		"resources":                   resources,
	}
}
