package ies

import "rrc/utils"

// ServingCellConfig-lte-NeighCellsCRS-Assumptions-r17 ::= ENUMERATED
type ServingcellconfigLteNeighcellscrsAssumptionsR17 struct {
	Value utils.ENUMERATED
}

const (
	ServingcellconfigLteNeighcellscrsAssumptionsR17EnumeratedNothing = iota
	ServingcellconfigLteNeighcellscrsAssumptionsR17EnumeratedFalse
)
