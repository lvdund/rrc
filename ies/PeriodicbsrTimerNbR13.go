package ies

import "rrc/utils"

// PeriodicBSR-Timer-NB-r13 ::= ENUMERATED
type PeriodicbsrTimerNbR13 struct {
	Value utils.ENUMERATED
}

const (
	PeriodicbsrTimerNbR13EnumeratedNothing = iota
	PeriodicbsrTimerNbR13EnumeratedPp2
	PeriodicbsrTimerNbR13EnumeratedPp4
	PeriodicbsrTimerNbR13EnumeratedPp8
	PeriodicbsrTimerNbR13EnumeratedPp16
	PeriodicbsrTimerNbR13EnumeratedPp64
	PeriodicbsrTimerNbR13EnumeratedPp128
	PeriodicbsrTimerNbR13EnumeratedInfinity
	PeriodicbsrTimerNbR13EnumeratedSpare
)
