package identity

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/hash/mimc"
)

type RandomNumGeneratorCircuit struct {
	PreImage frontend.Variable
	Hash     frontend.Variable
	Zero     frontend.Variable
}

func (circuit *RandomNumGeneratorCircuit) Define(api frontend.API) error {
	mimc, _ := mimc.NewMiMC("seed", api)

	mimc.Write(circuit.PreImage)
	api.AssertIsEqual(circuit.Hash, mimc.Sum())

	oof := api.FromBinary(api.ToBinary(mimc.Sum(), 64))

	const int64Max int = 18446744073709551615 / 2
	api.AssertIsLessOrEqual(oof, int64Max)
	api.AssertIsLessOrEqual(-int64Max, oof)

	const MAX int = 72

	api.AssertIsEqual(circuit.Zero, 0)

	randInt := api.Select((oof%api.Add(0, MAX) == circuit.Zero))

	return nil
}

func (num frontend.Variable) floor() frontend.Variable {
	return 0
}
