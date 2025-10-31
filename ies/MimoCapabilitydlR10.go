package ies

import "rrc/utils"

// MIMO-CapabilityDL-r10 ::= ENUMERATED
type MimoCapabilitydlR10 struct {
	Value utils.ENUMERATED
}

const (
	MimoCapabilitydlR10EnumeratedNothing = iota
	MimoCapabilitydlR10EnumeratedTwolayers
	MimoCapabilitydlR10EnumeratedFourlayers
	MimoCapabilitydlR10EnumeratedEightlayers
)
