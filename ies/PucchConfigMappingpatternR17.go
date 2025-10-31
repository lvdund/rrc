package ies

import "rrc/utils"

// PUCCH-Config-mappingPattern-r17 ::= ENUMERATED
type PucchConfigMappingpatternR17 struct {
	Value utils.ENUMERATED
}

const (
	PucchConfigMappingpatternR17EnumeratedNothing = iota
	PucchConfigMappingpatternR17EnumeratedCyclicmapping
	PucchConfigMappingpatternR17EnumeratedSequentialmapping
)
