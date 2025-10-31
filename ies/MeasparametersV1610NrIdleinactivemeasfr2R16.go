package ies

import "rrc/utils"

// MeasParameters-v1610-nr-IdleInactiveMeasFR2-r16 ::= ENUMERATED
type MeasparametersV1610NrIdleinactivemeasfr2R16 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1610NrIdleinactivemeasfr2R16EnumeratedNothing = iota
	MeasparametersV1610NrIdleinactivemeasfr2R16EnumeratedSupported
)
