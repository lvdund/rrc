package ies

// WhiteCellListNR-r16 ::= SEQUENCE OF PhysCellIdNR-r15
// SIZE (1..maxCellWhiteNR-r16)
type WhitecelllistnrR16 struct {
	Value []PhyscellidnrR15 `lb:1,ub:maxCellWhiteNRR16`
}
