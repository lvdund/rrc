package ies

import "rrc/utils"

// RRCSystemInfoRequest-IEs ::= SEQUENCE
type Rrcsysteminforequest struct {
	RequestedSiList utils.BITSTRING `lb:maxSIMessage,ub:maxSIMessage`
	Spare           utils.BITSTRING `lb:12,ub:12`
}
