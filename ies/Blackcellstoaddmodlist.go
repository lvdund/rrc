package ies

import "rrc/utils"

// BlackCellsToAddModList ::= SEQUENCE OF BlackCellsToAddMod
// SIZE (1..maxCellMeas)
type Blackcellstoaddmodlist struct {
	Value []Blackcellstoaddmod
}
