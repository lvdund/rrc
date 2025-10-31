package ies

import "rrc/utils"

// MeasParameters-v1610-nr-IdleInactiveMeasFR1-r16 ::= ENUMERATED
type MeasparametersV1610NrIdleinactivemeasfr1R16 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1610NrIdleinactivemeasfr1R16EnumeratedNothing = iota
	MeasparametersV1610NrIdleinactivemeasfr1R16EnumeratedSupported
)
