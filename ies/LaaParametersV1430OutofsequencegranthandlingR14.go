package ies

import "rrc/utils"

// LAA-Parameters-v1430-outOfSequenceGrantHandling-r14 ::= ENUMERATED
type LaaParametersV1430OutofsequencegranthandlingR14 struct {
	Value utils.ENUMERATED
}

const (
	LaaParametersV1430OutofsequencegranthandlingR14EnumeratedNothing = iota
	LaaParametersV1430OutofsequencegranthandlingR14EnumeratedSupported
)
