package ies

import "rrc/utils"

// SL-RLC-ChannelID-r17 ::= utils.INTEGER (1..maxSL-LCID-r16)
type SlRlcChannelidR17 struct {
	Value utils.INTEGER `lb:0,ub:maxSLLcidR16`
}
