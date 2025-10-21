package ies

import "rrc/utils"

// RetxBSR-Timer-NB-r13 ::= ENUMERATED
type RetxbsrTimerNbR13 struct {
	Value utils.ENUMERATED
}

const (
	RetxbsrTimerNbR13Pp4      = 0
	RetxbsrTimerNbR13Pp16     = 1
	RetxbsrTimerNbR13Pp64     = 2
	RetxbsrTimerNbR13Pp128    = 3
	RetxbsrTimerNbR13Pp256    = 4
	RetxbsrTimerNbR13Pp512    = 5
	RetxbsrTimerNbR13Infinity = 6
	RetxbsrTimerNbR13Spare    = 7
)
