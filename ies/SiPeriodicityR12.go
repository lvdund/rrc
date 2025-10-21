package ies

import "rrc/utils"

// SI-Periodicity-r12 ::= ENUMERATED
type SiPeriodicityR12 struct {
	Value utils.ENUMERATED
}

const (
	SiPeriodicityR12Rf8   = 0
	SiPeriodicityR12Rf16  = 1
	SiPeriodicityR12Rf32  = 2
	SiPeriodicityR12Rf64  = 3
	SiPeriodicityR12Rf128 = 4
	SiPeriodicityR12Rf256 = 5
	SiPeriodicityR12Rf512 = 6
)
