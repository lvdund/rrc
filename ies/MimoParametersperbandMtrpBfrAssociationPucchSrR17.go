package ies

import "rrc/utils"

// MIMO-ParametersPerBand-mTRP-BFR-association-PUCCH-SR-r17 ::= ENUMERATED
type MimoParametersperbandMtrpBfrAssociationPucchSrR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandMtrpBfrAssociationPucchSrR17EnumeratedNothing = iota
	MimoParametersperbandMtrpBfrAssociationPucchSrR17EnumeratedSupported
)
