package ies

import "rrc/utils"

// FR2-2-AccessParamsPerBand-r17-reduced-BeamSwitchTiming-FR2-2-r17 ::= ENUMERATED
type Fr22AccessparamsperbandR17ReducedBeamswitchtimingFr22R17 struct {
	Value utils.ENUMERATED
}

const (
	Fr22AccessparamsperbandR17ReducedBeamswitchtimingFr22R17EnumeratedNothing = iota
	Fr22AccessparamsperbandR17ReducedBeamswitchtimingFr22R17EnumeratedSupported
)
