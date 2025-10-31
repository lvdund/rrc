package ies

import "rrc/utils"

// RACH-ConfigGeneric-powerRampingStep ::= ENUMERATED
type RachConfiggenericPowerrampingstep struct {
	Value utils.ENUMERATED
}

const (
	RachConfiggenericPowerrampingstepEnumeratedNothing = iota
	RachConfiggenericPowerrampingstepEnumeratedDb0
	RachConfiggenericPowerrampingstepEnumeratedDb2
	RachConfiggenericPowerrampingstepEnumeratedDb4
	RachConfiggenericPowerrampingstepEnumeratedDb6
)
