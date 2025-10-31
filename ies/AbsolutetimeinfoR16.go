package ies

import "rrc/utils"

// AbsoluteTimeInfo-r16 ::= BIT STRING (SIZE (48))
type AbsolutetimeinfoR16 struct {
	Value utils.BITSTRING `lb:48,ub:48`
}
