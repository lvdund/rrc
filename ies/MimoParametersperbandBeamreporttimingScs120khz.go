package ies

import "rrc/utils"

// MIMO-ParametersPerBand-beamReportTiming-scs-120kHz ::= ENUMERATED
type MimoParametersperbandBeamreporttimingScs120khz struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandBeamreporttimingScs120khzEnumeratedNothing = iota
	MimoParametersperbandBeamreporttimingScs120khzEnumeratedSym14
	MimoParametersperbandBeamreporttimingScs120khzEnumeratedSym28
	MimoParametersperbandBeamreporttimingScs120khzEnumeratedSym56
)
