package ies

import "rrc/utils"

// BandParameters-v1530-qcl-TypeC-Operation-r15 ::= ENUMERATED
type BandparametersV1530QclTypecOperationR15 struct {
	Value utils.ENUMERATED
}

const (
	BandparametersV1530QclTypecOperationR15EnumeratedNothing = iota
	BandparametersV1530QclTypecOperationR15EnumeratedSupported
)
