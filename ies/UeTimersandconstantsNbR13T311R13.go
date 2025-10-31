package ies

import "rrc/utils"

// UE-TimersAndConstants-NB-r13-t311-r13 ::= ENUMERATED
type UeTimersandconstantsNbR13T311R13 struct {
	Value utils.ENUMERATED
}

const (
	UeTimersandconstantsNbR13T311R13EnumeratedNothing = iota
	UeTimersandconstantsNbR13T311R13EnumeratedMs1000
	UeTimersandconstantsNbR13T311R13EnumeratedMs3000
	UeTimersandconstantsNbR13T311R13EnumeratedMs5000
	UeTimersandconstantsNbR13T311R13EnumeratedMs10000
	UeTimersandconstantsNbR13T311R13EnumeratedMs15000
	UeTimersandconstantsNbR13T311R13EnumeratedMs20000
	UeTimersandconstantsNbR13T311R13EnumeratedMs30000
)
