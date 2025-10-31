package ies

import "rrc/utils"

// PRACH-ParametersCE-r13-prach-HoppingConfig-r13 ::= ENUMERATED
type PrachParametersceR13PrachHoppingconfigR13 struct {
	Value utils.ENUMERATED
}

const (
	PrachParametersceR13PrachHoppingconfigR13EnumeratedNothing = iota
	PrachParametersceR13PrachHoppingconfigR13EnumeratedOn
	PrachParametersceR13PrachHoppingconfigR13EnumeratedOff
)
