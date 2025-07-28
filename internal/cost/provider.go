package cost

import (
	"strings"

	"github.com/whittaker555/preflight/internal/models"
)

// DetectProvider inspects a Terraform plan and returns the dominant cloud provider.
func DetectProvider(plan models.TerraformPlan) string {
	counts := map[string]int{}
	for _, rc := range plan.ResourceChanges {
		switch {
		case strings.HasPrefix(rc.Type, "aws_"):
			counts["aws"]++
		case strings.HasPrefix(rc.Type, "google_"):
			counts["gcp"]++
		case strings.HasPrefix(rc.Type, "azurerm_"):
			counts["azure"]++
		}
	}

	detected := ""
	max := 0
	for p, c := range counts {
		if c > max {
			max = c
			detected = p
		}
	}

	if detected == "" {
		return "unknown"
	}
	return detected
}
