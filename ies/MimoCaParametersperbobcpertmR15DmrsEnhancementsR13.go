package ies

import "rrc/utils"

// MIMO-CA-ParametersPerBoBCPerTM-r15-dmrs-Enhancements-r13 ::= ENUMERATED
type MimoCaParametersperbobcpertmR15DmrsEnhancementsR13 struct {
	Value utils.ENUMERATED
}

const (
	MimoCaParametersperbobcpertmR15DmrsEnhancementsR13EnumeratedNothing = iota
	MimoCaParametersperbobcpertmR15DmrsEnhancementsR13EnumeratedDifferent
)
