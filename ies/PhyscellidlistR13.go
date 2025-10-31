package ies

// PhysCellIdList-r13 ::= SEQUENCE OF PhysCellId
// SIZE (1.. maxSL-DiscCells-r13)
type PhyscellidlistR13 struct {
	Value []Physcellid `lb:1,ub:maxSLDisccellsR13`
}
