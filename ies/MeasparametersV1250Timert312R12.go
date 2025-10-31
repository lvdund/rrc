package ies

import "rrc/utils"

// MeasParameters-v1250-timerT312-r12 ::= ENUMERATED
type MeasparametersV1250Timert312R12 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1250Timert312R12EnumeratedNothing = iota
	MeasparametersV1250Timert312R12EnumeratedSupported
)
