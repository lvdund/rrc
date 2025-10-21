package ies

import "rrc/utils"

// DRX-Config-r13 ::= SEQUENCE
type DrxConfigR13 struct {
	OndurationtimerV1310        *utils.ENUMERATED
	DrxRetransmissiontimerV1310 *utils.ENUMERATED
	DrxUlretransmissiontimerR13 *utils.ENUMERATED
}
