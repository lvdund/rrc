package ies

import "rrc/utils"

// PowerRampingParameters-powerRampingStep ::= ENUMERATED
type PowerrampingparametersPowerrampingstep struct {
	Value utils.ENUMERATED
}

const (
	PowerrampingparametersPowerrampingstepEnumeratedNothing = iota
	PowerrampingparametersPowerrampingstepEnumeratedDb0
	PowerrampingparametersPowerrampingstepEnumeratedDb2
	PowerrampingparametersPowerrampingstepEnumeratedDb4
	PowerrampingparametersPowerrampingstepEnumeratedDb6
)
