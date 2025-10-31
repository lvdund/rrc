package ies

import "rrc/utils"

// MIMO-ParametersPerBand-beamReportTiming-v1710-scs-960kHz-r17 ::= ENUMERATED
type MimoParametersperbandBeamreporttimingV1710Scs960khzR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandBeamreporttimingV1710Scs960khzR17EnumeratedNothing = iota
	MimoParametersperbandBeamreporttimingV1710Scs960khzR17EnumeratedSym112
	MimoParametersperbandBeamreporttimingV1710Scs960khzR17EnumeratedSym224
	MimoParametersperbandBeamreporttimingV1710Scs960khzR17EnumeratedSym448
)
