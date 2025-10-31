package ies

import "rrc/utils"

// MobilityParameters-v1610-cho-TwoTriggerEvents-r16 ::= ENUMERATED
type MobilityparametersV1610ChoTwotriggereventsR16 struct {
	Value utils.ENUMERATED
}

const (
	MobilityparametersV1610ChoTwotriggereventsR16EnumeratedNothing = iota
	MobilityparametersV1610ChoTwotriggereventsR16EnumeratedSupported
)
