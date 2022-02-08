package snark

import (
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
)

func TestFraction(t testing.T) {
	frac := FractionFromFloat(0.25)
	if !frac.numerator.Equal(new(fr.Element).SetInt64(1)) {
		t.Fatalf("numerator was not 1, got %d", frac.numerator.ToBigIntRegular(new(big.Int)))
	}
	if !frac.denominator.Equal(new(fr.Element).SetInt64(4)) {
		t.Fatalf("numerator was not 1, got %d", frac.denominator.ToBigIntRegular(new(big.Int)))
	}
}

func TestLCM(t testing.T) {
	lcm := LCM(*new(big.Int).SetInt64(4), *new(big.Int).SetInt64(6))
	if lcm.Cmp(new(big.Int).SetInt64(12)) != 0 {
		t.Fatalf("lcm was not 12, got %d", lcm.Int64())
	}
}

func TestAdd(t testing.T) {
	frac := NewFraction(1, 10)
	frac2 := NewFraction(2, 3)
	expectedResult := NewFraction(23, 30)
	addResult := new(Fraction).Add(frac, frac2)

	if addResult.Cmp(expectedResult) != 0 {
		t.Fatalf("expected %T got %T", expectedResult, addResult)
	}
}

func TestFloor(t testing.T) {
	frac := NewFraction(1, 10)
	floorResult := new(Fraction).Floor(frac)
	expectedResult := NewFraction(0, 10)

	if floorResult.Cmp(expectedResult) != 0 {
		t.Fatalf("expected %T got %T", expectedResult, floorResult)
	}

	frac = NewFraction(11, 10)
	floorResult = new(Fraction).Floor(frac)
	expectedResult = NewFraction(10, 10)

	if floorResult.Cmp(expectedResult) != 0 {
		t.Fatalf("expected %T got %T", expectedResult, floorResult)
	}
}

func TestMul(t testing.T) {
	frac := NewFraction(3, 4)
	frac2 := NewFraction(4, 3)
	expectedResult := NewFraction(12, 12)

	mulResult := new(Fraction).Mul(frac, frac2)

	if mulResult.Cmp(expectedResult) != 0 {
		t.Fatalf("expected %T got %T", expectedResult, mulResult)
	}
}

func TestDiv(t testing.T) {
	frac := NewFraction(3, 4)
	frac2 := NewFraction(1, 2)
	expectedResult := NewFraction(6, 4)

	divResult := new(Fraction).Div(frac, frac2)

	if divResult.Cmp(expectedResult) != 0 {
		t.Fatalf("expected %T got %T", expectedResult, divResult)
	}
}

func TestCmp(t testing.T) {
	frac := NewFraction(3, 4)
	frac2 := NewFraction(1, 2)

	if frac.Cmp(frac2) != 1 {
		t.Fatalf("comparison was not 1 got %d",
			frac.Cmp(frac2),
		)
	}
	if frac2.Cmp(frac) != -1 {
		t.Fatalf("comparison was not -1 got %d",
			frac2.Cmp(frac),
		)
	}
	if frac.Cmp(frac) != 0 {
		t.Fatalf("comparison was not 0 got %d",
			frac.Cmp(frac),
		)
	}
}

func TestSub(t testing.T) {
	frac := NewFraction(3, 4)
	frac2 := NewFraction(1, 2)
	expectedResult := NewFraction(1, 4)

	subResult := new(Fraction).Sub(frac, frac2)

	if new(Fraction).Sub(frac, frac2) != expectedResult {
		t.Fatalf("expected %T got %T", expectedResult, subResult)
	}
}

func TestMod(t testing.T) {
	frac := NewFraction(5, 1)
	frac2 := NewFraction(2, 1)
	expectedResult := NewFraction(1, 1)

	modResult := new(Fraction).Mod(frac, frac2)

	if modResult.Simplify() != expectedResult {
		t.Fatalf("expected %T got %T", expectedResult, modResult)
	}
}

func TestString(t testing.T) {
	frac := NewFraction(5, 1)

	if frac.String() != "Fraction: 5/1" {
		t.Fatalf("Expected 'Fraction: 5/1' got %s", frac.String())
	}
}
