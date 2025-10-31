package ies

// NS-PmaxList-r10 ::= SEQUENCE OF NS-PmaxValue-r10
// SIZE (1..maxNS-Pmax-r10)
type NsPmaxlistR10 struct {
	Value []NsPmaxvalueR10 `lb:1,ub:maxNSPmaxR10`
}
