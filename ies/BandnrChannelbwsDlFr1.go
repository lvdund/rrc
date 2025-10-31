package ies

import "rrc/utils"

// BandNR-channelBWs-DL-fr1 ::= SEQUENCE
type BandnrChannelbwsDlFr1 struct {
	Scs15khz *utils.BITSTRING `lb:10,ub:10`
	Scs30khz *utils.BITSTRING `lb:10,ub:10`
	Scs60khz *utils.BITSTRING `lb:10,ub:10`
}
