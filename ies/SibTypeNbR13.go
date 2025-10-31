package ies

import "rrc/utils"

// SIB-Type-NB-r13 ::= ENUMERATED
type SibTypeNbR13 struct {
	Value utils.ENUMERATED
}

const (
	SibTypeNbR13EnumeratedNothing = iota
	SibTypeNbR13EnumeratedSibtype3_Nb_R13
	SibTypeNbR13EnumeratedSibtype4_Nb_R13
	SibTypeNbR13EnumeratedSibtype5_Nb_R13
	SibTypeNbR13EnumeratedSibtype14_Nb_R13
	SibTypeNbR13EnumeratedSibtype16_Nb_R13
	SibTypeNbR13EnumeratedSibtype15_Nb_R14
	SibTypeNbR13EnumeratedSibtype20_Nb_R14
	SibTypeNbR13EnumeratedSibtype22_Nb_R14
)
