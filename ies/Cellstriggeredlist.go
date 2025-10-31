package ies

// CellsTriggeredList ::= SEQUENCE OF CellsTriggeredList-Item
// SIZE (1..maxNrofCellMeas)
type Cellstriggeredlist struct {
	Value []CellstriggeredlistItem `lb:1,ub:maxNrofCellMeas`
}
