package ies

import "rrc/utils"

// MIMO-ParametersPerBand-sfn-DefaultDL-BeamSetup-r17 ::= ENUMERATED
type MimoParametersperbandSfnDefaultdlBeamsetupR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSfnDefaultdlBeamsetupR17EnumeratedNothing = iota
	MimoParametersperbandSfnDefaultdlBeamsetupR17EnumeratedSupported
)
