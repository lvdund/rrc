package ies

// PLMN-RAN-AreaCell ::= SEQUENCE
type PlmnRanAreacell struct {
	PlmnIdentity *PlmnIdentity
	RanAreacells []Cellidentity `lb:1,ub:32`
}
