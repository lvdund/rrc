package ies

import "rrc/utils"

// RetxBSR-Timer-NB-r13 ::= ENUMERATED
type RetxbsrTimerNbR13 struct {
	Value utils.ENUMERATED
}

const (
	RetxbsrTimerNbR13EnumeratedNothing = iota
	RetxbsrTimerNbR13EnumeratedPp4
	RetxbsrTimerNbR13EnumeratedPp16
	RetxbsrTimerNbR13EnumeratedPp64
	RetxbsrTimerNbR13EnumeratedPp128
	RetxbsrTimerNbR13EnumeratedPp256
	RetxbsrTimerNbR13EnumeratedPp512
	RetxbsrTimerNbR13EnumeratedInfinity
	RetxbsrTimerNbR13EnumeratedSpare
)
