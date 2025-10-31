package ies

import "rrc/utils"

// GWUS-Config-r16-commonSequence-r16 ::= ENUMERATED
type GwusConfigR16CommonsequenceR16 struct {
	Value utils.ENUMERATED
}

const (
	GwusConfigR16CommonsequenceR16EnumeratedNothing = iota
	GwusConfigR16CommonsequenceR16EnumeratedG0
	GwusConfigR16CommonsequenceR16EnumeratedG126
)
