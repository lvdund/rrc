package ies

import "rrc/utils"

// PUSCH-Config-mappingPattern-r17 ::= ENUMERATED
type PuschConfigMappingpatternR17 struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigMappingpatternR17EnumeratedNothing = iota
	PuschConfigMappingpatternR17EnumeratedCyclicmapping
	PuschConfigMappingpatternR17EnumeratedSequentialmapping
)
