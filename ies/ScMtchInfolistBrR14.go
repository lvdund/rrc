package ies

// SC-MTCH-InfoList-BR-r14 ::= SEQUENCE OF SC-MTCH-Info-BR-r14
// SIZE (0..maxSC-MTCH-BR-r14)
type ScMtchInfolistBrR14 struct {
	Value []ScMtchInfoBrR14 `lb:0,ub:maxSCMtchBrR14`
}
