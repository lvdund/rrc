package ies

import "rrc/utils"

// BandNR-channelBWs-UL-v1590-fr2 ::= SEQUENCE
type BandnrChannelbwsUlV1590Fr2 struct {
	Scs60khz  *utils.BITSTRING `lb:8,ub:8`
	Scs120khz *utils.BITSTRING `lb:8,ub:8`
}
