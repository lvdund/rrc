package ies

import "rrc/utils"

// TrackingAreaCodeNR-r15 ::= BIT STRING (SIZE (24))
type TrackingareacodenrR15 struct {
	Value utils.BITSTRING `lb:24,ub:24`
}
