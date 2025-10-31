package ies

import "rrc/utils"

// SSB-PositionQCL-RelationNR-r16 ::= ENUMERATED
type SsbPositionqclRelationnrR16 struct {
	Value utils.ENUMERATED
}

const (
	SsbPositionqclRelationnrR16EnumeratedNothing = iota
	SsbPositionqclRelationnrR16EnumeratedN1
	SsbPositionqclRelationnrR16EnumeratedN2
	SsbPositionqclRelationnrR16EnumeratedN4
	SsbPositionqclRelationnrR16EnumeratedN8
)
