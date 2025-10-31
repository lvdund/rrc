package ies

import "rrc/utils"

// NPDSCH-MultiTB-Config-NB-r16-multiTB-Config-r16 ::= ENUMERATED
type NpdschMultitbConfigNbR16MultitbConfigR16 struct {
	Value utils.ENUMERATED
}

const (
	NpdschMultitbConfigNbR16MultitbConfigR16EnumeratedNothing = iota
	NpdschMultitbConfigNbR16MultitbConfigR16EnumeratedInterleaved
	NpdschMultitbConfigNbR16MultitbConfigR16EnumeratedNoninterleaved
)
