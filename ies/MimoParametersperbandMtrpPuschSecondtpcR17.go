package ies

import "rrc/utils"

// MIMO-ParametersPerBand-mTRP-PUSCH-secondTPC-r17 ::= ENUMERATED
type MimoParametersperbandMtrpPuschSecondtpcR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandMtrpPuschSecondtpcR17EnumeratedNothing = iota
	MimoParametersperbandMtrpPuschSecondtpcR17EnumeratedSupported
)
