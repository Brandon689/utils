package utils

import "testing"

func TestMyFunction(t *testing.T) {
	num := []float64{1, 4}
	result := Average(num)
	if result != 2.50 {
		t.Errorf("Average({1,4}) = %f; want 2.50", result)
	}
}
