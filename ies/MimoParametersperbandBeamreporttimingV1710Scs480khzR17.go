package ies

import "rrc/utils"

// MIMO-ParametersPerBand-beamReportTiming-v1710-scs-480kHz-r17 ::= ENUMERATED
type MimoParametersperbandBeamreporttimingV1710Scs480khzR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandBeamreporttimingV1710Scs480khzR17EnumeratedNothing = iota
	MimoParametersperbandBeamreporttimingV1710Scs480khzR17EnumeratedSym56
	MimoParametersperbandBeamreporttimingV1710Scs480khzR17EnumeratedSym112
	MimoParametersperbandBeamreporttimingV1710Scs480khzR17EnumeratedSym224
)
