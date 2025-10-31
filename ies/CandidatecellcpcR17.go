package ies

// CandidateCellCPC-r17 ::= SEQUENCE
type CandidatecellcpcR17 struct {
	SsbfrequencyR17      ArfcnValuenr
	CandidatecelllistR17 []Physcellid `lb:1,ub:maxNrofCondCellsR16`
}
