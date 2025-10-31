package ies

// UAC-BarringPerPLMN-List ::= SEQUENCE OF UAC-BarringPerPLMN
// SIZE (1.. maxPLMN)
type UacBarringperplmnList struct {
	Value []UacBarringperplmn `lb:1,ub:maxPLMN`
}
