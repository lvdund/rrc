package ies

import "rrc/utils"

// MIMO-ParametersPerBand-lowPAPR-DMRS-PUSCHwithPrecoding-r16 ::= ENUMERATED
type MimoParametersperbandLowpaprDmrsPuschwithprecodingR16 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandLowpaprDmrsPuschwithprecodingR16EnumeratedNothing = iota
	MimoParametersperbandLowpaprDmrsPuschwithprecodingR16EnumeratedSupported
)
