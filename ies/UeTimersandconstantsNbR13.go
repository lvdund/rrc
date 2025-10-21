package ies

import "rrc/utils"

// UE-TimersAndConstants-NB-r13 ::= SEQUENCE
// Extensible
type UeTimersandconstantsNbR13 struct {
	T300R13 utils.ENUMERATED
	T301R13 utils.ENUMERATED
	T310R13 utils.ENUMERATED
	N310R13 utils.ENUMERATED
	T311R13 utils.ENUMERATED
	N311R13 utils.ENUMERATED
}
