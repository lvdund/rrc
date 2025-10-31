package ies

import "rrc/utils"

// RLC-Parameters-extendedT-PollRetransmit-r16 ::= ENUMERATED
type RlcParametersExtendedtPollretransmitR16 struct {
	Value utils.ENUMERATED
}

const (
	RlcParametersExtendedtPollretransmitR16EnumeratedNothing = iota
	RlcParametersExtendedtPollretransmitR16EnumeratedSupported
)
