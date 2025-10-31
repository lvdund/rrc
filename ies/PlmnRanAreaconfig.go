package ies

// PLMN-RAN-AreaConfig ::= SEQUENCE
type PlmnRanAreaconfig struct {
	PlmnIdentity *PlmnIdentity
	RanArea      []RanAreaconfig `lb:1,ub:16`
}
