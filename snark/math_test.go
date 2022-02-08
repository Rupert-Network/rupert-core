package snark

import (
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
)

func TestFraction(t *testing.T) {
	frac := FractionFromFloat(0.25)
	if !frac.numerator.Equal(new(fr.Element).SetInt64(1)) {
		t.Fatalf("numerator was not 1, got %d", frac.numerator.ToBigIntRegular(new(big.Int)))
	}
	if !frac.denominator.Equal(new(fr.Element).SetInt64(4)) {
		t.Fatalf("numerator was not 1, got %d", frac.denominator.ToBigIntRegular(new(big.Int)))
	}
}

func TestLCM(t *testing.T) {
	lcm := LCM(*new(big.Int).SetInt64(4), *new(big.Int).SetInt64(6))
	if lcm.Cmp(new(big.Int).SetInt64(12)) != 0 {
		t.Fatalf("lcm was not 12, got %d", lcm.Int64())
	}
}

func TestAdd(t *testing.T) {
	frac := Fraction{
		*new(fr.Element).SetInt64(1),
		*new(fr.Element).SetInt64(10),
	}
	frac2 := Fraction{*new(fr.Element).SetInt64(2), *new(fr.Element).SetInt64(3)}

	frac3 := new(Fraction).Add(&frac, &frac2)

	expectedFrac := Fraction{*new(fr.Element).SetInt64(23), *new(fr.Element).SetInt64(30)}

	if !(frac3.numerator.Equal(&expectedFrac.numerator) && frac3.denominator.Equal(&expectedFrac.denominator)) {
		t.Fatalf("frac was not 23/30 got %d/%d",
			frac3.numerator.ToBigIntRegular(new(big.Int)),
			frac3.denominator.ToBigIntRegular(new(big.Int)),
		)
	}
}

func TestFloor(t *testing.T) {
	frac := Fraction{
		*new(fr.Element).SetInt64(1),
		*new(fr.Element).SetInt64(10),
	}

	flooredFrac := new(Fraction).Floor(&frac)

	if flooredFrac.numerator.Cmp(new(fr.Element).SetInt64(0)) != 0 {
		t.Fatalf("numerator was not 0 got %d",
			frac.numerator.ToBigIntRegular(new(big.Int)),
		)
	}

	frac = Fraction{
		*new(fr.Element).SetInt64(11),
		*new(fr.Element).SetInt64(10),
	}

	flooredFrac = new(Fraction).Floor(&frac)

	if flooredFrac.numerator.Cmp(new(fr.Element).SetInt64(10)) != 0 {
		t.Fatalf("numerator was not 10 got %d",
			frac.numerator.ToBigIntRegular(new(big.Int)),
		)
	}
}

func TestMul(t *testing.T) {
	frac := Fraction{
		*new(fr.Element).SetInt64(3),
		*new(fr.Element).SetInt64(4),
	}
	frac2 := Fraction{
		*new(fr.Element).SetInt64(4),
		*new(fr.Element).SetInt64(3),
	}
	mulFrac := new(Fraction).Mul(&frac, &frac2)

	if mulFrac.numerator.Cmp(new(fr.Element).SetInt64(12)) != 0 {
		t.Fatalf("numerator was not 12 got %d",
			mulFrac.numerator.ToBigIntRegular(new(big.Int)),
		)
	}
	if mulFrac.denominator.Cmp(new(fr.Element).SetInt64(12)) != 0 {
		t.Fatalf("numerator was not 12 got %d",
			mulFrac.numerator.ToBigIntRegular(new(big.Int)),
		)
	}
}

func TestDiv(t *testing.T) {
	frac := Fraction{
		*new(fr.Element).SetInt64(3),
		*new(fr.Element).SetInt64(4),
	}
	frac2 := Fraction{
		*new(fr.Element).SetInt64(1),
		*new(fr.Element).SetInt64(2),
	}
	divFrac := new(Fraction).Div(&frac, &frac2)

	if divFrac.numerator.Cmp(new(fr.Element).SetInt64(6)) != 0 {
		t.Fatalf("numerator was not 6 got %d/%d",
			divFrac.numerator.ToBigIntRegular(new(big.Int)),
			divFrac.denominator.ToBigIntRegular(new(big.Int)),
		)
	}
	if divFrac.denominator.Cmp(new(fr.Element).SetInt64(4)) != 0 {
		t.Fatalf("denominator was not 9 got %d/%d",
			divFrac.numerator.ToBigIntRegular(new(big.Int)),
			divFrac.denominator.ToBigIntRegular(new(big.Int)),
		)
	}
}

func TestCmp(t *testing.T) {
	frac := Fraction{
		*new(fr.Element).SetInt64(3),
		*new(fr.Element).SetInt64(4),
	}
	frac2 := Fraction{
		*new(fr.Element).SetInt64(1),
		*new(fr.Element).SetInt64(2),
	}
	if frac.Cmp(&frac2) != 1 {
		t.Fatalf("comparison was not 1 got %d",
			frac.Cmp(&frac2),
		)
	}
	if frac2.Cmp(&frac) != -1 {
		t.Fatalf("comparison was not -1 got %d",
			frac2.Cmp(&frac),
		)
	}
	if frac.Cmp(&frac) != 0 {
		t.Fatalf("comparison was not 0 got %d",
			frac.Cmp(&frac),
		)
	}
}

func TestSub(t *testing.T) {
	frac := Fraction{
		*new(fr.Element).SetInt64(3),
		*new(fr.Element).SetInt64(4),
	}
	frac2 := Fraction{
		*new(fr.Element).SetInt64(1),
		*new(fr.Element).SetInt64(2),
	}

	expectedResut := Fraction{
		*new(fr.Element).SetInt64(1),
		*new(fr.Element).SetInt64(4),
	}

	subResult := new(Fraction).Sub(&frac, &frac2)

	if *new(Fraction).Sub(&frac, &frac2) != expectedResut {
		t.Fatalf("expected Fraction %d/%d got Fraction %d/%d",
			expectedResut.numerator.ToBigIntRegular(new(big.Int)),
			expectedResut.denominator.ToBigIntRegular(new(big.Int)),
			subResult.numerator.ToBigIntRegular(new(big.Int)),
			subResult.denominator.ToBigIntRegular(new(big.Int)),
		)
	}
}
