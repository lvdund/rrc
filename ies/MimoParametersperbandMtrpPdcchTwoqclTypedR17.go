package ies

import "rrc/utils"

// MIMO-ParametersPerBand-mTRP-PDCCH-TwoQCL-TypeD-r17 ::= ENUMERATED
type MimoParametersperbandMtrpPdcchTwoqclTypedR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandMtrpPdcchTwoqclTypedR17EnumeratedNothing = iota
	MimoParametersperbandMtrpPdcchTwoqclTypedR17EnumeratedSupported
)
