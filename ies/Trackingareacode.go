package ies

import "rrc/utils"

// TrackingAreaCode ::= BIT STRING (SIZE (24))
type Trackingareacode struct {
	Value utils.BITSTRING `lb:24,ub:24`
}
