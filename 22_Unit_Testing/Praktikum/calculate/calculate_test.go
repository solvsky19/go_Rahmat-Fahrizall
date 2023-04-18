package calculate

import "testing"

func TestAddition(t *testing.T) {
	if Addition(2, 1) != 3 {
		t.Error("Expected 1 (+) 2 to equal 3")
	}
	if Addition(-2, -1) != -3 {
		t.Error("Expected -1 (+) -2 to equal -3")
	}
}

func TestSubstraction(t *testing.T) {
	if Substraction(2, 3) != -1 {
		t.Error("Expected 2 (-) 3 to equal -1")
	}
	if Substraction(3, 1) != 2 {
		t.Error("Expected 3 (-) 1 to equal 2")
	}
}

func TestMultiplication(t *testing.T) {
	if Multiplication(2, 2) != 4 {
		t.Error("Expected 2 (*) 2 to equal 4")
	}
	if Multiplication(3, 3) != 9 {
		t.Error("Expected 3 (*) 3 to equal 9")
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
