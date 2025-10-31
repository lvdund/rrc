package ies

import "rrc/utils"

// MIMO-ParametersPerBand-beamCorrespondenceCSI-RS-based-r16 ::= ENUMERATED
type MimoParametersperbandBeamcorrespondencecsiRsBasedR16 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandBeamcorrespondencecsiRsBasedR16EnumeratedNothing = iota
	MimoParametersperbandBeamcorrespondencecsiRsBasedR16EnumeratedSupported
)
