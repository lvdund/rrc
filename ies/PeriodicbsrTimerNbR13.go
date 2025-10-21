package ies

import "rrc/utils"

// PeriodicBSR-Timer-NB-r13 ::= ENUMERATED
type PeriodicbsrTimerNbR13 struct {
	Value utils.ENUMERATED
}

const (
	PeriodicbsrTimerNbR13Pp2      = 0
	PeriodicbsrTimerNbR13Pp4      = 1
	PeriodicbsrTimerNbR13Pp8      = 2
	PeriodicbsrTimerNbR13Pp16     = 3
	PeriodicbsrTimerNbR13Pp64     = 4
	PeriodicbsrTimerNbR13Pp128    = 5
	PeriodicbsrTimerNbR13Infinity = 6
	PeriodicbsrTimerNbR13Spare    = 7
)
