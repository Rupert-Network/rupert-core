package main

import (
	"github.com/consensys/gnark/frontend"
)

// Identity Creation
// User Bob and Validator Alice both precomit hashes to use for random selection
// User Bob provides a list of iris segment artifacts
// User Bob encrypts all said iris segments
// Validator alice validates the proof and posts the proof on the chain

type IrisArtifact struct {
	PosX frontend.Variable
	PosY frontend.Variable
}

type IDCreateCircuit struct {
	IrisArtifacts []IrisArtifact
}

func (circuit *IDCreateCircuit) Define(api frontend.API) error {

}
