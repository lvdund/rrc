package ies

import "rrc/utils"

// MobilityParameters-v1610-cho-r16 ::= ENUMERATED
type MobilityparametersV1610ChoR16 struct {
	Value utils.ENUMERATED
}

const (
	MobilityparametersV1610ChoR16EnumeratedNothing = iota
	MobilityparametersV1610ChoR16EnumeratedSupported
)
