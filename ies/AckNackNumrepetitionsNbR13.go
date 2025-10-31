package ies

import "rrc/utils"

// ACK-NACK-NumRepetitions-NB-r13 ::= ENUMERATED
type AckNackNumrepetitionsNbR13 struct {
	Value utils.ENUMERATED
}

const (
	AckNackNumrepetitionsNbR13EnumeratedNothing = iota
	AckNackNumrepetitionsNbR13EnumeratedR1
	AckNackNumrepetitionsNbR13EnumeratedR2
	AckNackNumrepetitionsNbR13EnumeratedR4
	AckNackNumrepetitionsNbR13EnumeratedR8
	AckNackNumrepetitionsNbR13EnumeratedR16
	AckNackNumrepetitionsNbR13EnumeratedR32
	AckNackNumrepetitionsNbR13EnumeratedR64
	AckNackNumrepetitionsNbR13EnumeratedR128
)
