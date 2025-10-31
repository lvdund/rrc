package ies

import "rrc/utils"

// MobilityParameters-v1610-cho-Failure-r16 ::= ENUMERATED
type MobilityparametersV1610ChoFailureR16 struct {
	Value utils.ENUMERATED
}

const (
	MobilityparametersV1610ChoFailureR16EnumeratedNothing = iota
	MobilityparametersV1610ChoFailureR16EnumeratedSupported
)
