package ies

// PMCH-InfoListExt-r12 ::= SEQUENCE OF PMCH-InfoExt-r12
// SIZE (0..maxPMCH-PerMBSFN)
type PmchInfolistextR12 struct {
	Value []PmchInfoextR12 `lb:0,ub:maxPMCHPermbsfn`
}
