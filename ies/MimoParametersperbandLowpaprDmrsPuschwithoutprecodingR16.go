package ies

import "rrc/utils"

// MIMO-ParametersPerBand-lowPAPR-DMRS-PUSCHwithoutPrecoding-r16 ::= ENUMERATED
type MimoParametersperbandLowpaprDmrsPuschwithoutprecodingR16 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandLowpaprDmrsPuschwithoutprecodingR16EnumeratedNothing = iota
	MimoParametersperbandLowpaprDmrsPuschwithoutprecodingR16EnumeratedSupported
)
