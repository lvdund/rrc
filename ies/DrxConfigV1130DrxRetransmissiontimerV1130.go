package ies

import "rrc/utils"

// DRX-Config-v1130-drx-RetransmissionTimer-v1130 ::= ENUMERATED
type DrxConfigV1130DrxRetransmissiontimerV1130 struct {
	Value utils.ENUMERATED
}

const (
	DrxConfigV1130DrxRetransmissiontimerV1130EnumeratedNothing = iota
	DrxConfigV1130DrxRetransmissiontimerV1130EnumeratedPsf0_V1130
)
