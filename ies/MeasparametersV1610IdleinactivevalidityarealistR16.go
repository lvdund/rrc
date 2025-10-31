package ies

import "rrc/utils"

// MeasParameters-v1610-idleInactiveValidityAreaList-r16 ::= ENUMERATED
type MeasparametersV1610IdleinactivevalidityarealistR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1610IdleinactivevalidityarealistR16EnumeratedNothing = iota
	MeasparametersV1610IdleinactivevalidityarealistR16EnumeratedSupported
)
