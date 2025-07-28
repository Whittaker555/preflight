package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

const samplePlan = `{
  "resource_changes": [
    {"address":"aws_instance.example","type":"aws_instance","name":"example","change":{"actions":["create"]}},
    {"address":"aws_s3_bucket.data","type":"aws_s3_bucket","name":"data","change":{"actions":["create"]}}
  ]
}`

func TestAnalysePlanHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/analyse", AnalysePlan)

	req := httptest.NewRequest(http.MethodPost, "/analyse", strings.NewReader(samplePlan))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var resp struct {
		Provider       string  `json:"provider"`
		TotalResources int     `json:"total_resources"`
		TotalCost      float64 `json:"total_monthly_cost_estimate"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}
	if resp.Provider != "aws" {
		t.Errorf("expected provider aws, got %s", resp.Provider)
	}
	if resp.TotalResources != 2 {
		t.Errorf("expected 2 resources, got %d", resp.TotalResources)
	}
	if resp.TotalCost != 30.0 {
		t.Errorf("expected cost 30.0, got %v", resp.TotalCost)
	}
}
