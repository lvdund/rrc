package ies

import "rrc/utils"

// MIMO-ParametersPerBand-mTRP-PUCCH-MAC-CE-r17 ::= ENUMERATED
type MimoParametersperbandMtrpPucchMacCeR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandMtrpPucchMacCeR17EnumeratedNothing = iota
	MimoParametersperbandMtrpPucchMacCeR17EnumeratedSupported
)
