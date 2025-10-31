package ies

import "rrc/utils"

// GapConfig-refServCellIndicator ::= ENUMERATED
type GapconfigRefservcellindicator struct {
	Value utils.ENUMERATED
}

const (
	GapconfigRefservcellindicatorEnumeratedNothing = iota
	GapconfigRefservcellindicatorEnumeratedPcell
	GapconfigRefservcellindicatorEnumeratedPscell
	GapconfigRefservcellindicatorEnumeratedMcg_Fr2
)
