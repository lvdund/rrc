package ies

import "rrc/utils"

// DRX-Config-v1130 ::= SEQUENCE
type DrxConfigV1130 struct {
	DrxRetransmissiontimerV1130  *utils.ENUMERATED
	LongdrxCyclestartoffsetV1130 *interface{}
	ShortdrxCycleV1130           *utils.ENUMERATED
}
