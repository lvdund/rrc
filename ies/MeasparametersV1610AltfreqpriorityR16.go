package ies

import "rrc/utils"

// MeasParameters-v1610-altFreqPriority-r16 ::= ENUMERATED
type MeasparametersV1610AltfreqpriorityR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1610AltfreqpriorityR16EnumeratedNothing = iota
	MeasparametersV1610AltfreqpriorityR16EnumeratedSupported
)
