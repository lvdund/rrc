package ies

import "rrc/utils"

// MIMO-ParametersPerBand-mTRP-BFR-PUCCH-SR-perCG-r17 ::= ENUMERATED
type MimoParametersperbandMtrpBfrPucchSrPercgR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandMtrpBfrPucchSrPercgR17EnumeratedNothing = iota
	MimoParametersperbandMtrpBfrPucchSrPercgR17EnumeratedN1
	MimoParametersperbandMtrpBfrPucchSrPercgR17EnumeratedN2
)
