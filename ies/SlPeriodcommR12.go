package ies

import "rrc/utils"

// SL-PeriodComm-r12 ::= ENUMERATED
type SlPeriodcommR12 struct {
	Value utils.ENUMERATED
}

const (
	SlPeriodcommR12EnumeratedNothing = iota
	SlPeriodcommR12EnumeratedSf40
	SlPeriodcommR12EnumeratedSf60
	SlPeriodcommR12EnumeratedSf70
	SlPeriodcommR12EnumeratedSf80
	SlPeriodcommR12EnumeratedSf120
	SlPeriodcommR12EnumeratedSf140
	SlPeriodcommR12EnumeratedSf160
	SlPeriodcommR12EnumeratedSf240
	SlPeriodcommR12EnumeratedSf280
	SlPeriodcommR12EnumeratedSf320
	SlPeriodcommR12EnumeratedSpare6
	SlPeriodcommR12EnumeratedSpare5
	SlPeriodcommR12EnumeratedSpare4
	SlPeriodcommR12EnumeratedSpare3
	SlPeriodcommR12EnumeratedSpare2
	SlPeriodcommR12EnumeratedSpare
)
