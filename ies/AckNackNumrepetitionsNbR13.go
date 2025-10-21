package ies

import "rrc/utils"

// ACK-NACK-NumRepetitions-NB-r13 ::= ENUMERATED
type AckNackNumrepetitionsNbR13 struct {
	Value utils.ENUMERATED
}

const (
	AckNackNumrepetitionsNbR13R1   = 0
	AckNackNumrepetitionsNbR13R2   = 1
	AckNackNumrepetitionsNbR13R4   = 2
	AckNackNumrepetitionsNbR13R8   = 3
	AckNackNumrepetitionsNbR13R16  = 4
	AckNackNumrepetitionsNbR13R32  = 5
	AckNackNumrepetitionsNbR13R64  = 6
	AckNackNumrepetitionsNbR13R128 = 7
)
