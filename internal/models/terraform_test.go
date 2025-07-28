package models

import (
	"encoding/json"
	"testing"
)

const samplePlan = `{
  "resource_changes": [
    {
      "address": "aws_instance.example",
      "type": "aws_instance",
      "name": "example",
      "change": {"actions": ["create"]}
    },
    {
      "address": "aws_s3_bucket.data",
      "type": "aws_s3_bucket",
      "name": "data",
      "change": {"actions": ["create"]}
    }
  ]
}`

func TestTerraformPlanUnmarshal(t *testing.T) {
	var p TerraformPlan
	if err := json.Unmarshal([]byte(samplePlan), &p); err != nil {
		t.Fatalf("failed to unmarshal plan: %v", err)
	}
	if len(p.ResourceChanges) != 2 {
		t.Fatalf("expected 2 resource changes, got %d", len(p.ResourceChanges))
	}
	if p.ResourceChanges[0].Address != "aws_instance.example" {
		t.Errorf("unexpected first address: %s", p.ResourceChanges[0].Address)
	}
}
