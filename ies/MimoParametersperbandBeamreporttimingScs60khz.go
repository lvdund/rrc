package ies

import "rrc/utils"

// MIMO-ParametersPerBand-beamReportTiming-scs-60kHz ::= ENUMERATED
type MimoParametersperbandBeamreporttimingScs60khz struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandBeamreporttimingScs60khzEnumeratedNothing = iota
	MimoParametersperbandBeamreporttimingScs60khzEnumeratedSym8
	MimoParametersperbandBeamreporttimingScs60khzEnumeratedSym14
	MimoParametersperbandBeamreporttimingScs60khzEnumeratedSym28
)
