package ies

import "rrc/utils"

// MeasParameters-v1610-measGapPatterns-NRonly-ENDC-r16 ::= ENUMERATED
type MeasparametersV1610MeasgappatternsNronlyEndcR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1610MeasgappatternsNronlyEndcR16EnumeratedNothing = iota
	MeasparametersV1610MeasgappatternsNronlyEndcR16EnumeratedSupported
)
