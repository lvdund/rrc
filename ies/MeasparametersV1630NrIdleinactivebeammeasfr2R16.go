package ies

import "rrc/utils"

// MeasParameters-v1630-nr-IdleInactiveBeamMeasFR2-r16 ::= ENUMERATED
type MeasparametersV1630NrIdleinactivebeammeasfr2R16 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1630NrIdleinactivebeammeasfr2R16EnumeratedNothing = iota
	MeasparametersV1630NrIdleinactivebeammeasfr2R16EnumeratedSupported
)
