package ies

// CandidateCellInfo-r17 ::= SEQUENCE
type CandidatecellinfoR17 struct {
	SsbfrequencyR17  ArfcnValuenr
	CandidatelistR17 []CandidatecellR17 `lb:1,ub:maxNrofCondCellsR16`
}
