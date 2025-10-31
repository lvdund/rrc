package ies

import "rrc/utils"

// GWUS-ResourceConfig-r16-resourceMappingPattern-r16 ::= CHOICE
const (
	GwusResourceconfigR16ResourcemappingpatternR16ChoiceNothing = iota
	GwusResourceconfigR16ResourcemappingpatternR16ChoiceResourcelocationwithwus
	GwusResourceconfigR16ResourcemappingpatternR16ChoiceResourcelocationwithoutwus
)

type GwusResourceconfigR16ResourcemappingpatternR16 struct {
	Choice                     uint64
	Resourcelocationwithwus    *utils.ENUMERATED
	Resourcelocationwithoutwus *utils.ENUMERATED
}
