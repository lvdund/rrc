package ies

import "rrc/utils"

// MeasParameters-v1630-ce-MeasRSS-DedicatedSameRBs-r16 ::= ENUMERATED
type MeasparametersV1630CeMeasrssDedicatedsamerbsR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1630CeMeasrssDedicatedsamerbsR16EnumeratedNothing = iota
	MeasparametersV1630CeMeasrssDedicatedsamerbsR16EnumeratedSupported
)
