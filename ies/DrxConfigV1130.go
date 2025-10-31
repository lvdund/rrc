package ies

// DRX-Config-v1130 ::= SEQUENCE
type DrxConfigV1130 struct {
	DrxRetransmissiontimerV1130  *DrxConfigV1130DrxRetransmissiontimerV1130
	LongdrxCyclestartoffsetV1130 *DrxConfigV1130LongdrxCyclestartoffsetV1130
	ShortdrxCycleV1130           *DrxConfigV1130ShortdrxCycleV1130
}
