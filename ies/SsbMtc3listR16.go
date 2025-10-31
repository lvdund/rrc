package ies

// SSB-MTC3List-r16 ::= SEQUENCE OF SSB-MTC3-r16
// SIZE (1..4)
type SsbMtc3listR16 struct {
	Value []SsbMtc3R16 `lb:1,ub:4`
}
