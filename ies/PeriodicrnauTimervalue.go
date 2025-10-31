package ies

import "rrc/utils"

// PeriodicRNAU-TimerValue ::= ENUMERATED
type PeriodicrnauTimervalue struct {
	Value utils.ENUMERATED
}

const (
	PeriodicrnauTimervalueEnumeratedNothing = iota
	PeriodicrnauTimervalueEnumeratedMin5
	PeriodicrnauTimervalueEnumeratedMin10
	PeriodicrnauTimervalueEnumeratedMin20
	PeriodicrnauTimervalueEnumeratedMin30
	PeriodicrnauTimervalueEnumeratedMin60
	PeriodicrnauTimervalueEnumeratedMin120
	PeriodicrnauTimervalueEnumeratedMin360
	PeriodicrnauTimervalueEnumeratedMin720
)
