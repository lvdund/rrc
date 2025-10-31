package ies

import "rrc/utils"

// TrackingAreaCode ::= BIT STRING (SIZE (16))
type Trackingareacode struct {
	Value utils.BITSTRING `lb:16,ub:16`
}
