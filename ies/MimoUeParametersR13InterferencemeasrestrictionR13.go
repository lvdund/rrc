package ies

import "rrc/utils"

// MIMO-UE-Parameters-r13-interferenceMeasRestriction-r13 ::= ENUMERATED
type MimoUeParametersR13InterferencemeasrestrictionR13 struct {
	Value utils.ENUMERATED
}

const (
	MimoUeParametersR13InterferencemeasrestrictionR13EnumeratedNothing = iota
	MimoUeParametersR13InterferencemeasrestrictionR13EnumeratedSupported
)
