package ies

import "rrc/utils"

// SchedulingInfo-si-Periodicity ::= ENUMERATED
type SchedulinginfoSiPeriodicity struct {
	Value utils.ENUMERATED
}

const (
	SchedulinginfoSiPeriodicityEnumeratedNothing = iota
	SchedulinginfoSiPeriodicityEnumeratedRf8
	SchedulinginfoSiPeriodicityEnumeratedRf16
	SchedulinginfoSiPeriodicityEnumeratedRf32
	SchedulinginfoSiPeriodicityEnumeratedRf64
	SchedulinginfoSiPeriodicityEnumeratedRf128
	SchedulinginfoSiPeriodicityEnumeratedRf256
	SchedulinginfoSiPeriodicityEnumeratedRf512
)
