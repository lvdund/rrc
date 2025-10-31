package ies

import "rrc/utils"

// MeasParameters-v1250-extendedMaxMeasId-r12 ::= ENUMERATED
type MeasparametersV1250ExtendedmaxmeasidR12 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1250ExtendedmaxmeasidR12EnumeratedNothing = iota
	MeasparametersV1250ExtendedmaxmeasidR12EnumeratedSupported
)
