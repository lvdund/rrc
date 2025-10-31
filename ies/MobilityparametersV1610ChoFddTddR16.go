package ies

import "rrc/utils"

// MobilityParameters-v1610-cho-FDD-TDD-r16 ::= ENUMERATED
type MobilityparametersV1610ChoFddTddR16 struct {
	Value utils.ENUMERATED
}

const (
	MobilityparametersV1610ChoFddTddR16EnumeratedNothing = iota
	MobilityparametersV1610ChoFddTddR16EnumeratedSupported
)
