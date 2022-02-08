package snark

import (
	"math"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
)

// Fraction ...
type Fraction struct {
	numerator   fr.Element
	denominator fr.Element
}

// NewFraction ...
func NewFraction(numerator, denominator int64) *Fraction {
	return new(Fraction).SetParts(numerator, denominator)
}

// String ...
func (frac *Fraction) String() string {
	return "Fraction: " + frac.numerator.String() + "/" + frac.denominator.String()
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

// SetParts ...
func (frac *Fraction) SetParts(num int64, denom int64) *Fraction {
	frac.numerator.SetInt64(num)
	frac.denominator.SetInt64(denom)

	return frac
}

// SetBigInt ...
func (frac *Fraction) SetBigInt(x big.Int) *Fraction {
	frac.numerator.SetBigInt(&x)
	frac.denominator.SetInt64(1)

	return frac
}

// Simplify ...
func (frac *Fraction) Simplify() *Fraction {
	a := frac.numerator.ToBigIntRegular(new(big.Int))
	b := frac.denominator.ToBigIntRegular(new(big.Int))

	gcd := new(big.Int).GCD(nil, nil, a, b)

	if gcd != new(big.Int).SetInt64(0) {
		newNumerator := new(big.Int).Div(a, gcd)
		newDenominator := new(big.Int).Div(b, gcd)

		return Fraction{*frac.numerator.SetBigInt(newNumerator), *frac.denominator.SetBigInt(newDenominator)}
	}

	return frac
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

// Floor ...
func (frac *Fraction) Floor(x *Fraction) *Fraction {
	if x.numerator.Cmp(&x.denominator) == -1 {
		frac.numerator = *new(fr.Element).SetInt64(0)
		frac.denominator = x.denominator

		return frac
	}

	_, m := new(big.Int).DivMod(
		x.numerator.ToBigIntRegular(new(big.Int)),
		x.denominator.ToBigIntRegular(new(big.Int)),
		new(big.Int),
	)

	frac.numerator = *new(fr.Element).Sub(
		&x.numerator,
		new(fr.Element).SetBigInt(m),
	)
	frac.denominator = x.denominator

	return frac
}

// Mul ...
func (frac *Fraction) Mul(x, y *Fraction) *Fraction {
	frac.numerator = *new(fr.Element).Mul(&y.numerator, &x.numerator)
	frac.denominator = *new(fr.Element).Mul(&y.denominator, &x.denominator)

	return frac
}

// Div ...
func (frac *Fraction) Div(x, y *Fraction) *Fraction {
	yNum := y.numerator
	y.numerator = y.denominator
	y.denominator = yNum

	frac.numerator = *new(fr.Element).Mul(&y.numerator, &x.numerator)
	frac.denominator = *new(fr.Element).Mul(&y.denominator, &x.denominator)

	return frac
}

// Cmp ...
func (frac *Fraction) Cmp(y *Fraction) int {
	if !frac.numerator.Equal(&y.numerator) {
		newDenom := LCM(
			*frac.denominator.ToBigIntRegular(new(big.Int)),
			*y.denominator.ToBigIntRegular(new(big.Int)),
		)

		multX := new(fr.Element).Div(new(fr.Element).SetBigInt(&newDenom), &frac.denominator)
		multY := new(fr.Element).Div(new(fr.Element).SetBigInt(&newDenom), &y.denominator)

		newXNum := new(fr.Element).Mul(multX, &frac.numerator)
		newYNum := new(fr.Element).Mul(multY, &y.numerator)

		return newXNum.Cmp(newYNum)
	}
	return frac.numerator.Cmp(&y.numerator)
}

// Sub ...
func (frac *Fraction) Sub(x, y *Fraction) *Fraction {
	if !x.numerator.Equal(&y.numerator) {
		newDenom := LCM(
			*x.denominator.ToBigIntRegular(new(big.Int)),
			*y.denominator.ToBigIntRegular(new(big.Int)),
		)

		multX := new(fr.Element).Div(new(fr.Element).SetBigInt(&newDenom), &x.denominator)
		multY := new(fr.Element).Div(new(fr.Element).SetBigInt(&newDenom), &y.denominator)

		newXNum := new(fr.Element).Mul(multX, &x.numerator)
		newYNum := new(fr.Element).Mul(multY, &y.numerator)

		frac.numerator = *new(fr.Element).Sub(newXNum, newYNum)
		frac.denominator = *new(fr.Element).SetBigInt(&newDenom)

	} else {

		newNumerator := new(fr.Element).Sub(&x.numerator, &y.numerator)

		frac.numerator = *newNumerator
		frac.denominator = x.denominator

	}
	return frac
}

// Mod ...
func (frac *Fraction) Mod(x, y *Fraction) *Fraction {
	divResult := new(Fraction).Div(x, y)

	return new(Fraction).Mul(
		new(Fraction).Floor(divResult),
		y,
	)
}
