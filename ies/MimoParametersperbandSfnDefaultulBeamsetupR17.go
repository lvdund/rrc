package ies

import "rrc/utils"

// MIMO-ParametersPerBand-sfn-DefaultUL-BeamSetup-r17 ::= ENUMERATED
type MimoParametersperbandSfnDefaultulBeamsetupR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSfnDefaultulBeamsetupR17EnumeratedNothing = iota
	MimoParametersperbandSfnDefaultulBeamsetupR17EnumeratedSupported
)
