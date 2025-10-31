package ies

// PLMN-RAN-AreaConfigList-r15 ::= SEQUENCE OF PLMN-RAN-AreaConfig-r15
// SIZE (1..maxPLMN-r15)
type PlmnRanAreaconfiglistR15 struct {
	Value []PlmnRanAreaconfigR15 `lb:1,ub:maxPLMNR15`
}
