package ies

import "rrc/utils"

// MIMO-ParametersPerBand-beamSwitchTiming-r16-scs-120kHz-r16 ::= ENUMERATED
type MimoParametersperbandBeamswitchtimingR16Scs120khzR16 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandBeamswitchtimingR16Scs120khzR16EnumeratedNothing = iota
	MimoParametersperbandBeamswitchtimingR16Scs120khzR16EnumeratedSym224
	MimoParametersperbandBeamswitchtimingR16Scs120khzR16EnumeratedSym336
)
