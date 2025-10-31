package ies

// PMCH-InfoList-r9 ::= SEQUENCE OF PMCH-Info-r9
// SIZE (0..maxPMCH-PerMBSFN)
type PmchInfolistR9 struct {
	Value []PmchInfoR9 `lb:0,ub:maxPMCHPermbsfn`
}
