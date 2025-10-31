package ies

import "rrc/utils"

// MeasParameters-v1310-extendedFreqPriorities-r13 ::= ENUMERATED
type MeasparametersV1310ExtendedfreqprioritiesR13 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1310ExtendedfreqprioritiesR13EnumeratedNothing = iota
	MeasparametersV1310ExtendedfreqprioritiesR13EnumeratedSupported
)
