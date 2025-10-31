package ies

import "rrc/utils"

// MIMO-ParametersPerBand-beamReportTiming-scs-30kHz ::= ENUMERATED
type MimoParametersperbandBeamreporttimingScs30khz struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandBeamreporttimingScs30khzEnumeratedNothing = iota
	MimoParametersperbandBeamreporttimingScs30khzEnumeratedSym4
	MimoParametersperbandBeamreporttimingScs30khzEnumeratedSym8
	MimoParametersperbandBeamreporttimingScs30khzEnumeratedSym14
	MimoParametersperbandBeamreporttimingScs30khzEnumeratedSym28
)
