package ies

import "rrc/utils"

// SC-MTCH-InfoList-NB-r14 ::= SEQUENCE OF SC-MTCH-Info-NB-r14
// SIZE (0.. maxSC-MTCH-NB-r14)
type ScMtchInfolistNbR14 struct {
	Value []ScMtchInfoNbR14
}
