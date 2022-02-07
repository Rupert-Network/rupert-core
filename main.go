package main

import (
	"fmt"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
)

type Circuit struct {
	X []frontend.Variable
	Y frontend.Variable
}

func (circuit *Circuit) Define(api frontend.API) error {
	for x := range circuit.X {
		api.AssertIsEqual(x, circuit.Y)
	}
	return nil
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("proof failed")
		}
	}()
	var cir Circuit
	ccs, _ := frontend.Compile(ecc.BN254, backend.GROTH16, &cir)

	pk, vk, _ := groth16.Setup(ccs)

	oof := []frontend.Variable{9, 9, 9}

	witness := Circuit{X: oof, Y: 9}

	proof, _ := groth16.Prove(ccs, pk, &witness)
	_ = groth16.Verify(proof, vk, &witness)
}
