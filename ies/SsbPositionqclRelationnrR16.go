package ies

import "rrc/utils"

// SSB-PositionQCL-RelationNR-r16 ::= ENUMERATED
type SsbPositionqclRelationnrR16 struct {
	Value utils.ENUMERATED
}

const (
	SsbPositionqclRelationnrR16N1 = 0
	SsbPositionqclRelationnrR16N2 = 1
	SsbPositionqclRelationnrR16N4 = 2
	SsbPositionqclRelationnrR16N8 = 3
)
