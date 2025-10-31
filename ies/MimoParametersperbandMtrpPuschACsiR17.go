package ies

import "rrc/utils"

// MIMO-ParametersPerBand-mTRP-PUSCH-A-CSI-r17 ::= ENUMERATED
type MimoParametersperbandMtrpPuschACsiR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandMtrpPuschACsiR17EnumeratedNothing = iota
	MimoParametersperbandMtrpPuschACsiR17EnumeratedSupported
)
