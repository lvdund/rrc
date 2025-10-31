package ies

// CondReconfigToAddModList-r16 ::= SEQUENCE OF CondReconfigToAddMod-r16
// SIZE (1.. maxNrofCondCells-r16)
type CondreconfigtoaddmodlistR16 struct {
	Value []CondreconfigtoaddmodR16 `lb:1,ub:maxNrofCondCellsR16`
}
