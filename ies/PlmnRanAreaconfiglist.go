package ies

// PLMN-RAN-AreaConfigList ::= SEQUENCE OF PLMN-RAN-AreaConfig
// SIZE (1..maxPLMNIdentities)
type PlmnRanAreaconfiglist struct {
	Value []PlmnRanAreaconfig `lb:1,ub:maxPLMNIdentities`
}
