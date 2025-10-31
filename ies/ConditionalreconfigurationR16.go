package ies

import "rrc/utils"

// ConditionalReconfiguration-r16 ::= SEQUENCE
// Extensible
type ConditionalreconfigurationR16 struct {
	AttemptcondreconfigR16      *utils.BOOLEAN
	CondreconfigtoremovelistR16 *CondreconfigtoremovelistR16
	CondreconfigtoaddmodlistR16 *CondreconfigtoaddmodlistR16
}
