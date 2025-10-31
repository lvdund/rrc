package ies

import "rrc/utils"

// BandNR-channelBWs-UL-v1590-fr1 ::= SEQUENCE
type BandnrChannelbwsUlV1590Fr1 struct {
	Scs15khz *utils.BITSTRING `lb:16,ub:16`
	Scs30khz *utils.BITSTRING `lb:16,ub:16`
	Scs60khz *utils.BITSTRING `lb:16,ub:16`
}
