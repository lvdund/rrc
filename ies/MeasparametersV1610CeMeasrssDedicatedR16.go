package ies

import "rrc/utils"

// MeasParameters-v1610-ce-MeasRSS-Dedicated-r16 ::= ENUMERATED
type MeasparametersV1610CeMeasrssDedicatedR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1610CeMeasrssDedicatedR16EnumeratedNothing = iota
	MeasparametersV1610CeMeasrssDedicatedR16EnumeratedSupported
)
