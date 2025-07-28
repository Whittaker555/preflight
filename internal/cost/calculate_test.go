package cost

import "testing"

func TestEstimateCostKnown(t *testing.T) {
	est := NewEstimator("aws")
	got := est.EstimateCost("aws_instance", 2)
	if got != 50.0 {
		t.Errorf("expected 50.0, got %v", got)
	}
}

func TestEstimateCostUnknown(t *testing.T) {
	est := NewEstimator("aws")
	got := est.EstimateCost("unknown_resource", 1)
	if got != 0.0 {
		t.Errorf("expected 0.0, got %v", got)
	}
}
