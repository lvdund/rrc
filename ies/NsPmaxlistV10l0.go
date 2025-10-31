package ies

// NS-PmaxList-v10l0 ::= SEQUENCE OF NS-PmaxValue-v10l0
// SIZE (1..maxNS-Pmax-r10)
type NsPmaxlistV10l0 struct {
	Value []NsPmaxvalueV10l0 `lb:1,ub:maxNSPmaxR10`
}
