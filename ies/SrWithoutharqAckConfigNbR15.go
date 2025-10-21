package ies

import "rrc/utils"

// SR-WithoutHARQ-ACK-Config-NB-r15 ::= CHOICE
type SrWithoutharqAckConfigNbR15 interface {
	isSrWithoutharqAckConfigNbR15()
}

type SrWithoutharqAckConfigNbR15Release struct {
	Value struct{}
}

func (*SrWithoutharqAckConfigNbR15Release) isSrWithoutharqAckConfigNbR15() {}

type SrWithoutharqAckConfigNbR15Setup struct {
	Value interface{}
}

func (*SrWithoutharqAckConfigNbR15Setup) isSrWithoutharqAckConfigNbR15() {}
