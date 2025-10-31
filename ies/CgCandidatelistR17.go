package ies

// CG-CandidateList-r17-IEs ::= SEQUENCE
type CgCandidatelistR17 struct {
	CgCandidatetoaddmodlistR17  *[]CgCandidateinfoR17   `lb:1,ub:maxNrofCondCellsR16`
	CgCandidatetoreleaselistR17 *[]CgCandidateinfoidR17 `lb:1,ub:maxNrofCondCellsR16`
	Noncriticalextension        *CgCandidatelistR17IesNoncriticalextension
}
