package snark

import (
	"math"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark/frontend"
)

// Fraction ...
type Fraction struct {
	numerator   fr.Element
	denominator fr.Element
}

// FractionFromFloat ...
func FractionFromFloat(num float64) Fraction {
	wholeNum := int(num)
	fracDecimal := num - float64(wholeNum)

	prec := 0

	for {
		if fracDecimal == float64(int(fracDecimal)) {
			break
		} else {
			prec++
			fracDecimal = fracDecimal * 10.0
		}
	}

	numerator := int64(prec*wholeNum + int(fracDecimal))
	denominator := int64(math.Pow(10, float64(prec)))

	unsimplified := Fraction{
		*new(fr.Element).SetInt64(numerator),
		*new(fr.Element).SetInt64(denominator),
	}

	return unsimplified.Simplify()

}

// Simplify ...
func (frac *Fraction) Simplify() Fraction {
	a := frac.numerator.ToBigIntRegular(new(big.Int))
	b := frac.denominator.ToBigIntRegular(new(big.Int))

	gcd := new(big.Int).GCD(nil, nil, a, b)

	if gcd != new(big.Int).SetInt64(0) {
		newNumerator := new(big.Int).Div(a, gcd)
		newDenominator := new(big.Int).Div(b, gcd)

		return Fraction{*frac.numerator.SetBigInt(newNumerator), *frac.denominator.SetBigInt(newDenominator)}
	}

	return *frac
}

// Floor ...
func Floor(api frontend.API, v frontend.Variable) frontend.Variable {

	return nil
}

// Modulus ...
func Modulus(v frontend.Variable, v2 frontend.Variable) frontend.Variable {
	// TODO: implemet
	// FIXME: needs floor implemented
	return 0
}

// LCM ...
func LCM(a, b big.Int) big.Int {
	aTimesB := new(big.Int).Mul(&a, &b)
	aBGCD := new(big.Int).GCD(nil, nil, &a, &b)

	return *new(big.Int).Div(aTimesB, aBGCD)
}

// Add ...
func (frac *Fraction) Add(x, y *Fraction) *Fraction {
	if !x.numerator.Equal(&y.numerator) {
		newDenom := LCM(
			*x.denominator.ToBigIntRegular(new(big.Int)),
			*y.denominator.ToBigIntRegular(new(big.Int)),
		)

		multX := new(fr.Element).Div(new(fr.Element).SetBigInt(&newDenom), &x.denominator)
		multY := new(fr.Element).Div(new(fr.Element).SetBigInt(&newDenom), &y.denominator)

		newXNum := new(fr.Element).Mul(multX, &x.numerator)
		newYNum := new(fr.Element).Mul(multY, &y.numerator)

		frac.numerator = *new(fr.Element).Add(newXNum, newYNum)
		frac.denominator = *new(fr.Element).SetBigInt(&newDenom)

	} else {

		newNumerator := new(fr.Element).Add(&x.numerator, &y.numerator)

		frac.numerator = *newNumerator
		frac.denominator = x.denominator

	}
	return frac
}
