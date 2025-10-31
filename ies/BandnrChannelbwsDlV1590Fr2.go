package ies

import "rrc/utils"

// BandNR-channelBWs-DL-v1590-fr2 ::= SEQUENCE
type BandnrChannelbwsDlV1590Fr2 struct {
	Scs60khz  *utils.BITSTRING `lb:8,ub:8`
	Scs120khz *utils.BITSTRING `lb:8,ub:8`
}
