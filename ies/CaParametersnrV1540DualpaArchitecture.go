package ies

import "rrc/utils"

// CA-ParametersNR-v1540-dualPA-Architecture ::= ENUMERATED
type CaParametersnrV1540DualpaArchitecture struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1540DualpaArchitectureEnumeratedNothing = iota
	CaParametersnrV1540DualpaArchitectureEnumeratedSupported
)
