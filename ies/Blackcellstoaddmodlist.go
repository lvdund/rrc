package ies

// BlackCellsToAddModList ::= SEQUENCE OF BlackCellsToAddMod
// SIZE (1..maxCellMeas)
type Blackcellstoaddmodlist struct {
	Value []Blackcellstoaddmod `lb:1,ub:maxCellMeas`
}
