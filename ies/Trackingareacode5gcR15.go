package ies

import "rrc/utils"

// TrackingAreaCode-5GC-r15 ::= BIT STRING (SIZE (24))
type Trackingareacode5gcR15 struct {
	Value utils.BITSTRING `lb:24,ub:24`
}
