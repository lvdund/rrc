package ies

// CellsTriggeredList ::= SEQUENCE OF CellsTriggeredList-Item
// SIZE (1..maxCellMeas)
type Cellstriggeredlist struct {
	Value []CellstriggeredlistItem `lb:1,ub:maxCellMeas`
}
