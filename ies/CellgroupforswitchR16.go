package ies

// CellGroupForSwitch-r16 ::= SEQUENCE OF ServCellIndex
// SIZE (1..16)
type CellgroupforswitchR16 struct {
	Value []Servcellindex `lb:1,ub:16`
}
