package ies

import "rrc/utils"

// BH-RLC-ChannelID-r16 ::= BIT STRING (SIZE (16))
type BhRlcChannelidR16 struct {
	Value utils.BITSTRING `lb:16,ub:16`
}
