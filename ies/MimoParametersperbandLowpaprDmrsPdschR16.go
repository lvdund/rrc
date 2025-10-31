package ies

import "rrc/utils"

// MIMO-ParametersPerBand-lowPAPR-DMRS-PDSCH-r16 ::= ENUMERATED
type MimoParametersperbandLowpaprDmrsPdschR16 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandLowpaprDmrsPdschR16EnumeratedNothing = iota
	MimoParametersperbandLowpaprDmrsPdschR16EnumeratedSupported
)
