package ies

import "rrc/utils"

// SIB-Type-NB-v1530 ::= ENUMERATED
type SibTypeNbV1530 struct {
	Value utils.ENUMERATED
}

const (
	SibTypeNbV1530Sibtype23NbR15 = 0
	SibTypeNbV1530Sibtype27NbR16 = 1
	SibTypeNbV1530Spare6         = 2
	SibTypeNbV1530Spare5         = 3
	SibTypeNbV1530Spare4         = 4
	SibTypeNbV1530Spare3         = 5
	SibTypeNbV1530Spare2         = 6
	SibTypeNbV1530Spare1         = 7
)
