package ies

import "rrc/utils"

// BandNR-channelBWs-DL-fr2 ::= SEQUENCE
type BandnrChannelbwsDlFr2 struct {
	Scs60khz  *utils.BITSTRING `lb:3,ub:3`
	Scs120khz *utils.BITSTRING `lb:3,ub:3`
}
