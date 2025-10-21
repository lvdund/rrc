package ies

import "rrc/utils"

// CellsTriggeredList ::= SEQUENCE OF CHOICE
// SIZE (1..maxCellMeas)
type Cellstriggeredlist struct {
	Value []Choice
}
