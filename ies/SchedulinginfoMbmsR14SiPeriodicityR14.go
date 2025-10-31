package ies

import "rrc/utils"

// SchedulingInfo-MBMS-r14-si-Periodicity-r14 ::= ENUMERATED
type SchedulinginfoMbmsR14SiPeriodicityR14 struct {
	Value utils.ENUMERATED
}

const (
	SchedulinginfoMbmsR14SiPeriodicityR14EnumeratedNothing = iota
	SchedulinginfoMbmsR14SiPeriodicityR14EnumeratedRf16
	SchedulinginfoMbmsR14SiPeriodicityR14EnumeratedRf32
	SchedulinginfoMbmsR14SiPeriodicityR14EnumeratedRf64
	SchedulinginfoMbmsR14SiPeriodicityR14EnumeratedRf128
	SchedulinginfoMbmsR14SiPeriodicityR14EnumeratedRf256
	SchedulinginfoMbmsR14SiPeriodicityR14EnumeratedRf512
)
