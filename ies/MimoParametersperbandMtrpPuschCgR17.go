package ies

import "rrc/utils"

// MIMO-ParametersPerBand-mTRP-PUSCH-CG-r17 ::= ENUMERATED
type MimoParametersperbandMtrpPuschCgR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandMtrpPuschCgR17EnumeratedNothing = iota
	MimoParametersperbandMtrpPuschCgR17EnumeratedSupported
)
