package ies

// AdditionalRACH-ConfigList-r17 ::= SEQUENCE OF AdditionalRACH-Config-r17
// SIZE (1..maxAdditionalRACH-r17)
type AdditionalrachConfiglistR17 struct {
	Value []AdditionalrachConfigR17 `lb:1,ub:maxAdditionalRACHR17`
}
