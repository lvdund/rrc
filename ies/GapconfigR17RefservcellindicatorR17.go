package ies

import "rrc/utils"

// GapConfig-r17-refServCellIndicator-r17 ::= ENUMERATED
type GapconfigR17RefservcellindicatorR17 struct {
	Value utils.ENUMERATED
}

const (
	GapconfigR17RefservcellindicatorR17EnumeratedNothing = iota
	GapconfigR17RefservcellindicatorR17EnumeratedPcell
	GapconfigR17RefservcellindicatorR17EnumeratedPscell
	GapconfigR17RefservcellindicatorR17EnumeratedMcg_Fr2
)
