package ies

import "rrc/utils"

// BandParametersTxSL-r14-v2x-eNB-Scheduled-r14 ::= ENUMERATED
type BandparameterstxslR14V2xEnbScheduledR14 struct {
	Value utils.ENUMERATED
}

const (
	BandparameterstxslR14V2xEnbScheduledR14EnumeratedNothing = iota
	BandparameterstxslR14V2xEnbScheduledR14EnumeratedSupported
)
