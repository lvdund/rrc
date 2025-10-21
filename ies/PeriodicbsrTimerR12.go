package ies

import "rrc/utils"

// PeriodicBSR-Timer-r12 ::= ENUMERATED
type PeriodicbsrTimerR12 struct {
	Value utils.ENUMERATED
}

const (
	PeriodicbsrTimerR12Sf5      = 0
	PeriodicbsrTimerR12Sf10     = 1
	PeriodicbsrTimerR12Sf16     = 2
	PeriodicbsrTimerR12Sf20     = 3
	PeriodicbsrTimerR12Sf32     = 4
	PeriodicbsrTimerR12Sf40     = 5
	PeriodicbsrTimerR12Sf64     = 6
	PeriodicbsrTimerR12Sf80     = 7
	PeriodicbsrTimerR12Sf128    = 8
	PeriodicbsrTimerR12Sf160    = 9
	PeriodicbsrTimerR12Sf320    = 10
	PeriodicbsrTimerR12Sf640    = 11
	PeriodicbsrTimerR12Sf1280   = 12
	PeriodicbsrTimerR12Sf2560   = 13
	PeriodicbsrTimerR12Infinity = 14
	PeriodicbsrTimerR12Spare1   = 15
)
