package ies

import "rrc/utils"

// MIMO-ParametersPerBand-beamCorrespondenceSSB-based-r16 ::= ENUMERATED
type MimoParametersperbandBeamcorrespondencessbBasedR16 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandBeamcorrespondencessbBasedR16EnumeratedNothing = iota
	MimoParametersperbandBeamcorrespondencessbBasedR16EnumeratedSupported
)
