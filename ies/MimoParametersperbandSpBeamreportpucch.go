package ies

import "rrc/utils"

// MIMO-ParametersPerBand-sp-BeamReportPUCCH ::= ENUMERATED
type MimoParametersperbandSpBeamreportpucch struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSpBeamreportpucchEnumeratedNothing = iota
	MimoParametersperbandSpBeamreportpucchEnumeratedSupported
)
