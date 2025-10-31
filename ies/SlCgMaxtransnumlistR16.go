package ies

// SL-CG-MaxTransNumList-r16 ::= SEQUENCE OF SL-CG-MaxTransNum-r16
// SIZE (1..8)
type SlCgMaxtransnumlistR16 struct {
	Value []SlCgMaxtransnumR16 `lb:1,ub:8`
}
