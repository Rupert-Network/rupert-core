package snark

import (
	"math"
	"math/big"

	"github.com/consensys/gnark/frontend"
)

// Fraction ...
type Fraction struct {
	numerator   int64
	denominator int64
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

	numerator := prec*wholeNum + int(fracDecimal)
	denominator := int64(math.Pow(10, float64(prec)))

	unsimplified := Fraction{int64(numerator), denominator}

	return unsimplified.simplify()

}

func (frac *Fraction) simplify() Fraction {
	a := new(big.Int).SetInt64(frac.numerator)
	b := new(big.Int).SetInt64(frac.denominator)

	gcd := new(big.Int).GCD(nil, nil, a, b)

	if gcd != new(big.Int).SetInt64(0) {
		newNumerator := new(big.Int).Div(a, gcd)
		newDenominator := new(big.Int).Div(b, gcd)

		return Fraction{newNumerator.Int64(), newDenominator.Int64()}
	}

	return *frac
}

// FrontendVariables ...
func (frac *Fraction) FrontendVariables() (frontend.Variable, frontend.Variable) {
	return frac.numerator, frac.denominator
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
