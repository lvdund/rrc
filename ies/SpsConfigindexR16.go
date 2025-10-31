package ies

import "rrc/utils"

// SPS-ConfigIndex-r16 ::= utils.INTEGER (0.. maxNrofSPS-Config-1-r16)
type SpsConfigindexR16 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofSPSConfig1R16`
}
