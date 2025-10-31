package ies

// PhysCellIdRangeNR-r16 ::= SEQUENCE
type PhyscellidrangenrR16 struct {
	Start PhyscellidnrR15
	Range *PhyscellidrangenrR16Range
}
