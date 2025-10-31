package ies

import "rrc/utils"

// MIMO-ParametersPerBand-lowPAPR-DMRS-PUCCH-r16 ::= ENUMERATED
type MimoParametersperbandLowpaprDmrsPucchR16 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandLowpaprDmrsPucchR16EnumeratedNothing = iota
	MimoParametersperbandLowpaprDmrsPucchR16EnumeratedSupported
)
