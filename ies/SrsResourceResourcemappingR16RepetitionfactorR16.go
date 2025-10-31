package ies

import "rrc/utils"

// SRS-Resource-resourceMapping-r16-repetitionFactor-r16 ::= ENUMERATED
type SrsResourceResourcemappingR16RepetitionfactorR16 struct {
	Value utils.ENUMERATED
}

const (
	SrsResourceResourcemappingR16RepetitionfactorR16EnumeratedNothing = iota
	SrsResourceResourcemappingR16RepetitionfactorR16EnumeratedN1
	SrsResourceResourcemappingR16RepetitionfactorR16EnumeratedN2
	SrsResourceResourcemappingR16RepetitionfactorR16EnumeratedN4
)
