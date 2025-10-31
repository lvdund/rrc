package ies

import "rrc/utils"

// MIMO-ParametersPerBand-sp-BeamReportPUSCH ::= ENUMERATED
type MimoParametersperbandSpBeamreportpusch struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSpBeamreportpuschEnumeratedNothing = iota
	MimoParametersperbandSpBeamreportpuschEnumeratedSupported
)
