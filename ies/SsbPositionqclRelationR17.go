package ies

import "rrc/utils"

// SSB-PositionQCL-Relation-r17 ::= ENUMERATED
type SsbPositionqclRelationR17 struct {
	Value utils.ENUMERATED
}

const (
	SsbPositionqclRelationR17EnumeratedNothing = iota
	SsbPositionqclRelationR17EnumeratedN32
	SsbPositionqclRelationR17EnumeratedN64
)
