package ies

// AltTTT-CellsToAddModList-r12 ::= SEQUENCE OF AltTTT-CellsToAddMod-r12
// SIZE (1..maxCellMeas)
type AlttttCellstoaddmodlistR12 struct {
	Value []AlttttCellstoaddmodR12 `lb:1,ub:maxCellMeas`
}
