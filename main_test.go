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

func TestMaxInt64(t *testing.T) {
	if MaxInt64(0, -1) != 0 {
		t.Errorf("Expected %v to be > %v", 0, -1)
	}

	if MaxInt64(-5, -4) != -4 {
		t.Errorf("Expected %v to be > %v", -4, -5)
	}
}

func TestAbsInt(t *testing.T) {
	if AbsInt(-3) != 3 {
		t.Errorf("Expected %v to be %v", AbsInt(-3), 3)
	}

	if AbsInt(-7-3) != 10 {
		t.Errorf("Expected %v to be %v", AbsInt(-7-3), 10)
	}
}

func TestMinInt(t *testing.T) {
	if MinInt(0, -1) != -1 {
		t.Errorf("Expected %v to be < %v", -1, 0)
	}

	if MinInt(-5, -4) != -5 {
		t.Errorf("Expected %v to be < %v", -5, -4)
	}
}

func TestMaxInt(t *testing.T) {
	if MaxInt(0, -1) != 0 {
		t.Errorf("Expected %v to be > %v", 0, -1)
	}

	if MaxInt(-5, -4) != -4 {
		t.Errorf("Expected %v to be > %v", -4, -5)
	}
}
