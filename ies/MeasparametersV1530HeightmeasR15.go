package ies

import "rrc/utils"

// MeasParameters-v1530-heightMeas-r15 ::= ENUMERATED
type MeasparametersV1530HeightmeasR15 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1530HeightmeasR15EnumeratedNothing = iota
	MeasparametersV1530HeightmeasR15EnumeratedSupported
)
