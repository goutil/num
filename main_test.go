package num

import "testing"

func TestMinInt64(t *testing.T) {
	if MinInt64(0, -1) != -1 {
		t.Errorf("Expected %v to be < %v", -1, 0)
	}

	if MinInt64(-5, -4) != -5 {
		t.Errorf("Expected %v to be < %v", -5, -4)
	}
}
