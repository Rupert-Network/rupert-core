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
	frac := Fraction{*new(fr.Element).SetInt64(1), *new(fr.Element).SetInt64(10)}
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
