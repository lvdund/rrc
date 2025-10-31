package ies

import "rrc/utils"

// MIMO-ParametersPerBand-beamSwitchTiming-r17-scs-480kHz-r17 ::= ENUMERATED
type MimoParametersperbandBeamswitchtimingR17Scs480khzR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandBeamswitchtimingR17Scs480khzR17EnumeratedNothing = iota
	MimoParametersperbandBeamswitchtimingR17Scs480khzR17EnumeratedSym896
	MimoParametersperbandBeamswitchtimingR17Scs480khzR17EnumeratedSym1344
)
