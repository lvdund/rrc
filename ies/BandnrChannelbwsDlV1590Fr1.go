package ies

import "rrc/utils"

// BandNR-channelBWs-DL-v1590-fr1 ::= SEQUENCE
type BandnrChannelbwsDlV1590Fr1 struct {
	Scs15khz *utils.BITSTRING `lb:16,ub:16`
	Scs30khz *utils.BITSTRING `lb:16,ub:16`
	Scs60khz *utils.BITSTRING `lb:16,ub:16`
}
