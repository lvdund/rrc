package ies

import "rrc/utils"

// ServingCellConfig-directionalCollisionHandling-r16 ::= ENUMERATED
type ServingcellconfigDirectionalcollisionhandlingR16 struct {
	Value utils.ENUMERATED
}

const (
	ServingcellconfigDirectionalcollisionhandlingR16EnumeratedNothing = iota
	ServingcellconfigDirectionalcollisionhandlingR16EnumeratedEnabled
)
