package ies

// SC-MTCH-InfoList-r13 ::= SEQUENCE OF SC-MTCH-Info-r13
// SIZE (0..maxSC-MTCH-r13)
type ScMtchInfolistR13 struct {
	Value []ScMtchInfoR13 `lb:0,ub:maxSCMtchR13`
}
