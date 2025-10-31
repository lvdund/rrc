package ies

import "rrc/utils"

// DownlinkHARQ-FeedbackDisabled-r17 ::= BIT STRING (SIZE (32))
type DownlinkharqFeedbackdisabledR17 struct {
	Value utils.BITSTRING `lb:32,ub:32`
}
