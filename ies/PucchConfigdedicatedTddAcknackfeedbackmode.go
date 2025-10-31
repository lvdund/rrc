package ies

import "rrc/utils"

// PUCCH-ConfigDedicated-tdd-AckNackFeedbackMode ::= ENUMERATED
type PucchConfigdedicatedTddAcknackfeedbackmode struct {
	Value utils.ENUMERATED
}

const (
	PucchConfigdedicatedTddAcknackfeedbackmodeEnumeratedNothing = iota
	PucchConfigdedicatedTddAcknackfeedbackmodeEnumeratedBundling
	PucchConfigdedicatedTddAcknackfeedbackmodeEnumeratedMultiplexing
)
