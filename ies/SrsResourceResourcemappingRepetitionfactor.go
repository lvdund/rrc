package ies

import "rrc/utils"

// SRS-Resource-resourceMapping-repetitionFactor ::= ENUMERATED
type SrsResourceResourcemappingRepetitionfactor struct {
	Value utils.ENUMERATED
}

const (
	SrsResourceResourcemappingRepetitionfactorEnumeratedNothing = iota
	SrsResourceResourcemappingRepetitionfactorEnumeratedN1
	SrsResourceResourcemappingRepetitionfactorEnumeratedN2
	SrsResourceResourcemappingRepetitionfactorEnumeratedN4
)
