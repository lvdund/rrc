package ies

import "rrc/utils"

// MIMO-ParametersPerBand-beamSwitchTiming-r16-scs-60kHz-r16 ::= ENUMERATED
type MimoParametersperbandBeamswitchtimingR16Scs60khzR16 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandBeamswitchtimingR16Scs60khzR16EnumeratedNothing = iota
	MimoParametersperbandBeamswitchtimingR16Scs60khzR16EnumeratedSym224
	MimoParametersperbandBeamswitchtimingR16Scs60khzR16EnumeratedSym336
)
