package ies

import "rrc/utils"

// PUCCH-ConfigDedicated-r13-tdd-AckNackFeedbackMode-r13 ::= ENUMERATED
type PucchConfigdedicatedR13TddAcknackfeedbackmodeR13 struct {
	Value utils.ENUMERATED
}

const (
	PucchConfigdedicatedR13TddAcknackfeedbackmodeR13EnumeratedNothing = iota
	PucchConfigdedicatedR13TddAcknackfeedbackmodeR13EnumeratedBundling
	PucchConfigdedicatedR13TddAcknackfeedbackmodeR13EnumeratedMultiplexing
)
