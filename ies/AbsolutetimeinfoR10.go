package ies

import "rrc/utils"

// AbsoluteTimeInfo-r10 ::= BIT STRING (SIZE (48))
type AbsolutetimeinfoR10 struct {
	Value utils.BITSTRING `lb:48,ub:48`
}
