package ies

// MIMO-ParametersPerBand-beamSwitchTiming ::= SEQUENCE
type MimoParametersperbandBeamswitchtiming struct {
	Scs60khz  *MimoParametersperbandBeamswitchtimingScs60khz
	Scs120khz *MimoParametersperbandBeamswitchtimingScs120khz
}
