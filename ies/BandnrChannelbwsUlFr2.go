package ies

import "rrc/utils"

// BandNR-channelBWs-UL-fr2 ::= SEQUENCE
type BandnrChannelbwsUlFr2 struct {
	Scs60khz  *utils.BITSTRING `lb:3,ub:3`
	Scs120khz *utils.BITSTRING `lb:3,ub:3`
}
