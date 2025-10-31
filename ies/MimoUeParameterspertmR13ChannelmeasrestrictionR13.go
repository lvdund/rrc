package ies

import "rrc/utils"

// MIMO-UE-ParametersPerTM-r13-channelMeasRestriction-r13 ::= ENUMERATED
type MimoUeParameterspertmR13ChannelmeasrestrictionR13 struct {
	Value utils.ENUMERATED
}

const (
	MimoUeParameterspertmR13ChannelmeasrestrictionR13EnumeratedNothing = iota
	MimoUeParameterspertmR13ChannelmeasrestrictionR13EnumeratedSupported
)
