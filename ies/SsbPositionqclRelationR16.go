package ies

import "rrc/utils"

// SSB-PositionQCL-Relation-r16 ::= ENUMERATED
type SsbPositionqclRelationR16 struct {
	Value utils.ENUMERATED
}

const (
	SsbPositionqclRelationR16EnumeratedNothing = iota
	SsbPositionqclRelationR16EnumeratedN1
	SsbPositionqclRelationR16EnumeratedN2
	SsbPositionqclRelationR16EnumeratedN4
	SsbPositionqclRelationR16EnumeratedN8
)
