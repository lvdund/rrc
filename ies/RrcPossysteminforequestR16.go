package ies

import "rrc/utils"

// RRC-PosSystemInfoRequest-r16-IEs ::= SEQUENCE
type RrcPossysteminforequestR16 struct {
	RequestedpossiList utils.BITSTRING `lb:maxSIMessage,ub:maxSIMessage`
	Spare              utils.BITSTRING `lb:11,ub:11`
}
