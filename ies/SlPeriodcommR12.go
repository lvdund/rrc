package ies

import "rrc/utils"

// SL-PeriodComm-r12 ::= ENUMERATED
type SlPeriodcommR12 struct {
	Value utils.ENUMERATED
}

const (
	SlPeriodcommR12Sf40   = 0
	SlPeriodcommR12Sf60   = 1
	SlPeriodcommR12Sf70   = 2
	SlPeriodcommR12Sf80   = 3
	SlPeriodcommR12Sf120  = 4
	SlPeriodcommR12Sf140  = 5
	SlPeriodcommR12Sf160  = 6
	SlPeriodcommR12Sf240  = 7
	SlPeriodcommR12Sf280  = 8
	SlPeriodcommR12Sf320  = 9
	SlPeriodcommR12Spare6 = 10
	SlPeriodcommR12Spare5 = 11
	SlPeriodcommR12Spare4 = 12
	SlPeriodcommR12Spare3 = 13
	SlPeriodcommR12Spare2 = 14
	SlPeriodcommR12Spare  = 15
)
