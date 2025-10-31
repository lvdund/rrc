package ies

import "rrc/utils"

// UE-TimersAndConstants-NB-r13-t310-r13 ::= ENUMERATED
type UeTimersandconstantsNbR13T310R13 struct {
	Value utils.ENUMERATED
}

const (
	UeTimersandconstantsNbR13T310R13EnumeratedNothing = iota
	UeTimersandconstantsNbR13T310R13EnumeratedMs0
	UeTimersandconstantsNbR13T310R13EnumeratedMs200
	UeTimersandconstantsNbR13T310R13EnumeratedMs500
	UeTimersandconstantsNbR13T310R13EnumeratedMs1000
	UeTimersandconstantsNbR13T310R13EnumeratedMs2000
	UeTimersandconstantsNbR13T310R13EnumeratedMs4000
	UeTimersandconstantsNbR13T310R13EnumeratedMs8000
)
