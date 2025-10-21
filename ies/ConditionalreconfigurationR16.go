package ies

import "rrc/utils"

// ConditionalReconfiguration-r16 ::= SEQUENCE
// Extensible
type ConditionalreconfigurationR16 struct {
	CondreconfigurationtoaddmodlistR16 *CondreconfigurationtoaddmodlistR16
	CondreconfigurationtoremovelistR16 *CondreconfigurationtoremovelistR16
	AttemptcondreconfR16               *utils.ENUMERATED
}
