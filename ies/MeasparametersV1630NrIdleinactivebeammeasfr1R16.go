package ies

import "rrc/utils"

// MeasParameters-v1630-nr-IdleInactiveBeamMeasFR1-r16 ::= ENUMERATED
type MeasparametersV1630NrIdleinactivebeammeasfr1R16 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1630NrIdleinactivebeammeasfr1R16EnumeratedNothing = iota
	MeasparametersV1630NrIdleinactivebeammeasfr1R16EnumeratedSupported
)
