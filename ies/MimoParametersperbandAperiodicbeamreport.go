package ies

import "rrc/utils"

// MIMO-ParametersPerBand-aperiodicBeamReport ::= ENUMERATED
type MimoParametersperbandAperiodicbeamreport struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandAperiodicbeamreportEnumeratedNothing = iota
	MimoParametersperbandAperiodicbeamreportEnumeratedSupported
)
