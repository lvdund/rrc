package ies

import "rrc/utils"

// GWUS-Config-NB-r16-commonSequence-r16 ::= ENUMERATED
type GwusConfigNbR16CommonsequenceR16 struct {
	Value utils.ENUMERATED
}

const (
	GwusConfigNbR16CommonsequenceR16EnumeratedNothing = iota
	GwusConfigNbR16CommonsequenceR16EnumeratedG0
	GwusConfigNbR16CommonsequenceR16EnumeratedG126
)
