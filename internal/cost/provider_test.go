package cost

import (
	"testing"

	"github.com/whittaker555/preflight/internal/models"
)

func TestDetectProviderAWS(t *testing.T) {
	plan := models.TerraformPlan{
		ResourceChanges: []models.ResourceChange{
			{Type: "aws_instance"},
			{Type: "aws_s3_bucket"},
		},
	}
	if got := DetectProvider(plan); got != "aws" {
		t.Errorf("expected aws, got %s", got)
	}
}

func TestDetectProviderGCP(t *testing.T) {
	plan := models.TerraformPlan{
		ResourceChanges: []models.ResourceChange{
			{Type: "google_compute_instance"},
		},
	}
	if got := DetectProvider(plan); got != "gcp" {
		t.Errorf("expected gcp, got %s", got)
	}
}
