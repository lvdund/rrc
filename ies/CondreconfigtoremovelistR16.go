package ies

// CondReconfigToRemoveList-r16 ::= SEQUENCE OF CondReconfigId-r16
// SIZE (1.. maxNrofCondCells-r16)
type CondreconfigtoremovelistR16 struct {
	Value []CondreconfigidR16 `lb:1,ub:maxNrofCondCellsR16`
}
