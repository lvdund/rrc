package ies

// WhiteCellsToAddModList-r13 ::= SEQUENCE OF WhiteCellsToAddMod-r13
// SIZE (1..maxCellMeas)
type WhitecellstoaddmodlistR13 struct {
	Value []WhitecellstoaddmodR13 `lb:1,ub:maxCellMeas`
}
