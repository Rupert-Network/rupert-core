package snark

import "testing"

func TestFraction(t *testing.T) {
	frac := FractionFromFloat(0.25)
	if frac.numerator != 1 {
		t.Fatalf("numerator was not 1, got %d", frac.numerator)
	}
	if frac.denominator != 4 {
		t.Fatalf("denominator was not 4, got %d", frac.denominator)
	}
}
