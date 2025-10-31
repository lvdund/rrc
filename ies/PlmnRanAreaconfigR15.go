package ies

// PLMN-RAN-AreaConfig-r15 ::= SEQUENCE
type PlmnRanAreaconfigR15 struct {
	PlmnIdentityR15 *PlmnIdentity
	RanAreaR15      []RanAreaconfigR15 `lb:1,ub:16`
}
