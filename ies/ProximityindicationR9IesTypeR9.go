package ies

import "rrc/utils"

// ProximityIndication-r9-IEs-type-r9 ::= ENUMERATED
type ProximityindicationR9IesTypeR9 struct {
	Value utils.ENUMERATED
}

const (
	ProximityindicationR9IesTypeR9EnumeratedNothing = iota
	ProximityindicationR9IesTypeR9EnumeratedEntering
	ProximityindicationR9IesTypeR9EnumeratedLeaving
)
