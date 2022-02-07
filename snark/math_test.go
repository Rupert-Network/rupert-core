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
