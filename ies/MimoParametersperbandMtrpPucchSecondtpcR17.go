package ies

import "rrc/utils"

// MIMO-ParametersPerBand-mTRP-PUCCH-SecondTPC-r17 ::= ENUMERATED
type MimoParametersperbandMtrpPucchSecondtpcR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandMtrpPucchSecondtpcR17EnumeratedNothing = iota
	MimoParametersperbandMtrpPucchSecondtpcR17EnumeratedSupported
)
