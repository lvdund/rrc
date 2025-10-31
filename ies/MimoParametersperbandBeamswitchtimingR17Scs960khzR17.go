package ies

import "rrc/utils"

// MIMO-ParametersPerBand-beamSwitchTiming-r17-scs-960kHz-r17 ::= ENUMERATED
type MimoParametersperbandBeamswitchtimingR17Scs960khzR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandBeamswitchtimingR17Scs960khzR17EnumeratedNothing = iota
	MimoParametersperbandBeamswitchtimingR17Scs960khzR17EnumeratedSym1792
	MimoParametersperbandBeamswitchtimingR17Scs960khzR17EnumeratedSym2688
)
