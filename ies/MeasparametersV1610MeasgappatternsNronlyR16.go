package ies

import "rrc/utils"

// MeasParameters-v1610-measGapPatterns-NRonly-r16 ::= ENUMERATED
type MeasparametersV1610MeasgappatternsNronlyR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1610MeasgappatternsNronlyR16EnumeratedNothing = iota
	MeasparametersV1610MeasgappatternsNronlyR16EnumeratedSupported
)
