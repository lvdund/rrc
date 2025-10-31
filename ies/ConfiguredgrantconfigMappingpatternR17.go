package ies

import "rrc/utils"

// ConfiguredGrantConfig-mappingPattern-r17 ::= ENUMERATED
type ConfiguredgrantconfigMappingpatternR17 struct {
	Value utils.ENUMERATED
}

const (
	ConfiguredgrantconfigMappingpatternR17EnumeratedNothing = iota
	ConfiguredgrantconfigMappingpatternR17EnumeratedCyclicmapping
	ConfiguredgrantconfigMappingpatternR17EnumeratedSequentialmapping
)
