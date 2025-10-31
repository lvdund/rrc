package ies

import "rrc/utils"

// SC-MTCH-Info-BR-r14-p-a-r14 ::= ENUMERATED
type ScMtchInfoBrR14PAR14 struct {
	Value utils.ENUMERATED
}

const (
	ScMtchInfoBrR14PAR14EnumeratedNothing = iota
	ScMtchInfoBrR14PAR14EnumeratedDb_6
	ScMtchInfoBrR14PAR14EnumeratedDb_4dot77
	ScMtchInfoBrR14PAR14EnumeratedDb_3
	ScMtchInfoBrR14PAR14EnumeratedDb_1dot77
	ScMtchInfoBrR14PAR14EnumeratedDb0
	ScMtchInfoBrR14PAR14EnumeratedDb1
	ScMtchInfoBrR14PAR14EnumeratedDb2
	ScMtchInfoBrR14PAR14EnumeratedDb3
)
