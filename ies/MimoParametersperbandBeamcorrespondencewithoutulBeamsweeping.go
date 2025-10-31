package ies

import "rrc/utils"

// MIMO-ParametersPerBand-beamCorrespondenceWithoutUL-BeamSweeping ::= ENUMERATED
type MimoParametersperbandBeamcorrespondencewithoutulBeamsweeping struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandBeamcorrespondencewithoutulBeamsweepingEnumeratedNothing = iota
	MimoParametersperbandBeamcorrespondencewithoutulBeamsweepingEnumeratedSupported
)
