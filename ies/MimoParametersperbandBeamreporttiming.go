package ies

// MIMO-ParametersPerBand-beamReportTiming ::= SEQUENCE
type MimoParametersperbandBeamreporttiming struct {
	Scs15khz  *MimoParametersperbandBeamreporttimingScs15khz
	Scs30khz  *MimoParametersperbandBeamreporttimingScs30khz
	Scs60khz  *MimoParametersperbandBeamreporttimingScs60khz
	Scs120khz *MimoParametersperbandBeamreporttimingScs120khz
}
