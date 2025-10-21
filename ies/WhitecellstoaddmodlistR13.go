package ies

import "rrc/utils"

// WhiteCellsToAddModList-r13 ::= SEQUENCE OF WhiteCellsToAddMod-r13
// SIZE (1..maxCellMeas)
type WhitecellstoaddmodlistR13 struct {
	Value utils.Sequence[WhitecellstoaddmodR13]
}
