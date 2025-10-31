package ies

import "rrc/utils"

// SI-Periodicity-r12 ::= ENUMERATED
type SiPeriodicityR12 struct {
	Value utils.ENUMERATED
}

const (
	SiPeriodicityR12EnumeratedNothing = iota
	SiPeriodicityR12EnumeratedRf8
	SiPeriodicityR12EnumeratedRf16
	SiPeriodicityR12EnumeratedRf32
	SiPeriodicityR12EnumeratedRf64
	SiPeriodicityR12EnumeratedRf128
	SiPeriodicityR12EnumeratedRf256
	SiPeriodicityR12EnumeratedRf512
)
