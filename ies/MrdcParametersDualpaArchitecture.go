package ies

import "rrc/utils"

// MRDC-Parameters-dualPA-Architecture ::= ENUMERATED
type MrdcParametersDualpaArchitecture struct {
	Value utils.ENUMERATED
}

const (
	MrdcParametersDualpaArchitectureEnumeratedNothing = iota
	MrdcParametersDualpaArchitectureEnumeratedSupported
)
