package ies

// PLMN-RAN-AreaCell-r15 ::= SEQUENCE
type PlmnRanAreacellR15 struct {
	PlmnIdentityR15 *PlmnIdentity
	RanAreacellsR15 []Cellidentity `lb:1,ub:32`
}
