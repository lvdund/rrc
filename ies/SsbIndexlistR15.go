package ies

// SSB-IndexList-r15 ::= SEQUENCE OF RS-IndexNR-r15
// SIZE (1..maxRS-Index-r15)
type SsbIndexlistR15 struct {
	Value []RsIndexnrR15 `lb:1,ub:maxRSIndexR15`
}
