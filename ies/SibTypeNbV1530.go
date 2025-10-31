package ies

import "rrc/utils"

// SIB-Type-NB-v1530 ::= ENUMERATED
type SibTypeNbV1530 struct {
	Value utils.ENUMERATED
}

const (
	SibTypeNbV1530EnumeratedNothing = iota
	SibTypeNbV1530EnumeratedSibtype23_Nb_R15
	SibTypeNbV1530EnumeratedSibtype27_Nb_R16
	SibTypeNbV1530EnumeratedSpare6
	SibTypeNbV1530EnumeratedSpare5
	SibTypeNbV1530EnumeratedSpare4
	SibTypeNbV1530EnumeratedSpare3
	SibTypeNbV1530EnumeratedSpare2
	SibTypeNbV1530EnumeratedSpare1
)
