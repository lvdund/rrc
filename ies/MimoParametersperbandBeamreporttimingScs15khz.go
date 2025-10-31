package ies

import "rrc/utils"

// MIMO-ParametersPerBand-beamReportTiming-scs-15kHz ::= ENUMERATED
type MimoParametersperbandBeamreporttimingScs15khz struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandBeamreporttimingScs15khzEnumeratedNothing = iota
	MimoParametersperbandBeamreporttimingScs15khzEnumeratedSym2
	MimoParametersperbandBeamreporttimingScs15khzEnumeratedSym4
	MimoParametersperbandBeamreporttimingScs15khzEnumeratedSym8
)
