package ies

import "rrc/utils"

// MIMO-ParametersPerBand-periodicBeamReport ::= ENUMERATED
type MimoParametersperbandPeriodicbeamreport struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandPeriodicbeamreportEnumeratedNothing = iota
	MimoParametersperbandPeriodicbeamreportEnumeratedSupported
)
