package ies

import "rrc/utils"

// MIMO-ParametersPerBand-mTRP-PDCCH-individual-r17 ::= ENUMERATED
type MimoParametersperbandMtrpPdcchIndividualR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandMtrpPdcchIndividualR17EnumeratedNothing = iota
	MimoParametersperbandMtrpPdcchIndividualR17EnumeratedSupported
)
