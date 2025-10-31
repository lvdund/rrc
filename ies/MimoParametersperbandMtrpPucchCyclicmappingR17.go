package ies

import "rrc/utils"

// MIMO-ParametersPerBand-mTRP-PUCCH-CyclicMapping-r17 ::= ENUMERATED
type MimoParametersperbandMtrpPucchCyclicmappingR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandMtrpPucchCyclicmappingR17EnumeratedNothing = iota
	MimoParametersperbandMtrpPucchCyclicmappingR17EnumeratedSupported
)
