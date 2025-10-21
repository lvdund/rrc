package ies

import "rrc/utils"

// RetxBSR-Timer-r12 ::= ENUMERATED
type RetxbsrTimerR12 struct {
	Value utils.ENUMERATED
}

const (
	RetxbsrTimerR12Sf320   = 0
	RetxbsrTimerR12Sf640   = 1
	RetxbsrTimerR12Sf1280  = 2
	RetxbsrTimerR12Sf2560  = 3
	RetxbsrTimerR12Sf5120  = 4
	RetxbsrTimerR12Sf10240 = 5
	RetxbsrTimerR12Spare2  = 6
	RetxbsrTimerR12Spare1  = 7
)
