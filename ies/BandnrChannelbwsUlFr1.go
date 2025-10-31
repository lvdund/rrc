package ies

import "rrc/utils"

// BandNR-channelBWs-UL-fr1 ::= SEQUENCE
type BandnrChannelbwsUlFr1 struct {
	Scs15khz *utils.BITSTRING `lb:10,ub:10`
	Scs30khz *utils.BITSTRING `lb:10,ub:10`
	Scs60khz *utils.BITSTRING `lb:10,ub:10`
}
