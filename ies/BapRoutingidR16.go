package ies

import "rrc/utils"

// BAP-RoutingID-r16 ::= SEQUENCE
type BapRoutingidR16 struct {
	BapAddressR16 utils.BITSTRING `lb:10,ub:10`
	BapPathidR16  utils.BITSTRING `lb:10,ub:10`
}
