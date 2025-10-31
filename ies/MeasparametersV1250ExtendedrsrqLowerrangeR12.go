package ies

import "rrc/utils"

// MeasParameters-v1250-extendedRSRQ-LowerRange-r12 ::= ENUMERATED
type MeasparametersV1250ExtendedrsrqLowerrangeR12 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1250ExtendedrsrqLowerrangeR12EnumeratedNothing = iota
	MeasparametersV1250ExtendedrsrqLowerrangeR12EnumeratedSupported
)
