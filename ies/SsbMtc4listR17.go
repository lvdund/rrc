package ies

// SSB-MTC4List-r17 ::= SEQUENCE OF SSB-MTC4-r17
// SIZE (1..3)
type SsbMtc4listR17 struct {
	Value []SsbMtc4R17 `lb:1,ub:3`
}
