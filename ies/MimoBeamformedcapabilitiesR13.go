package ies

import "rrc/utils"

// MIMO-BeamformedCapabilities-r13 ::= SEQUENCE
type MimoBeamformedcapabilitiesR13 struct {
	KMaxR13     utils.INTEGER    `lb:0,ub:8`
	NMaxlistR13 *utils.BITSTRING `lb:1,ub:7`
}
