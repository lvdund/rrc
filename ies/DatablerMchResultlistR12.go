package ies

// DataBLER-MCH-ResultList-r12 ::= SEQUENCE OF DataBLER-MCH-Result-r12
// SIZE (1.. maxPMCH-PerMBSFN)
type DatablerMchResultlistR12 struct {
	Value []DatablerMchResultR12 `lb:1,ub:maxPMCHPermbsfn`
}
