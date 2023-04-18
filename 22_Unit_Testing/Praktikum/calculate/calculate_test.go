package calculate

import "testing"

func TestAddition(t *testing.T) {
	if Addition(4, 2) != 6 {
		t.Error("Expected 4 (+) 2 to equal 6")
	}
	if Addition(-4, -2) != -6 {
		t.Error("Expected -4 (+) -2 to equal -6")
	}
}

func TestSubstraction(t *testing.T) {
	if Substraction(7, 3) != 4 {
		t.Error("Expected 7 (-) 3 to equal 4")
	}
	if Substraction(12, 6) != 6 {
		t.Error("Expected 12 (-) 6 to equal 6")
	}
}

func TestMultiplication(t *testing.T) {
	if Multiplication(3, 3) != 9 {
		t.Error("Expected 3 (*) 3 to equal 9")
	}
	if Multiplication(4, 2) != 8 {
		t.Error("Expected 4 (*) 2 to equal 8")
	}
}

func TestDivision(t *testing.T) {
	if Division(6, 3) != 2 {
		t.Error("Expected 6 (/) 3 to equal 2")
	}
	if Division(8, 2) != 4 {
		t.Error("Expected 8 (/) 2 to equal 4")
	}
}
