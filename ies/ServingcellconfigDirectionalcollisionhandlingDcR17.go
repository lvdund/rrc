package ies

import "rrc/utils"

// ServingCellConfig-directionalCollisionHandling-DC-r17 ::= ENUMERATED
type ServingcellconfigDirectionalcollisionhandlingDcR17 struct {
	Value utils.ENUMERATED
}

const (
	ServingcellconfigDirectionalcollisionhandlingDcR17EnumeratedNothing = iota
	ServingcellconfigDirectionalcollisionhandlingDcR17EnumeratedEnabled
)
