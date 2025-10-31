package ies

import "rrc/utils"

// PUSCH-ConfigDedicated-r13-pusch-HoppingConfig-r13 ::= ENUMERATED
type PuschConfigdedicatedR13PuschHoppingconfigR13 struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigdedicatedR13PuschHoppingconfigR13EnumeratedNothing = iota
	PuschConfigdedicatedR13PuschHoppingconfigR13EnumeratedOn
)
