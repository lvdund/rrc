package ies

import "rrc/utils"

// UE-TimersAndConstants-t311 ::= ENUMERATED
type UeTimersandconstantsT311 struct {
	Value utils.ENUMERATED
}

const (
	UeTimersandconstantsT311EnumeratedNothing = iota
	UeTimersandconstantsT311EnumeratedMs1000
	UeTimersandconstantsT311EnumeratedMs3000
	UeTimersandconstantsT311EnumeratedMs5000
	UeTimersandconstantsT311EnumeratedMs10000
	UeTimersandconstantsT311EnumeratedMs15000
	UeTimersandconstantsT311EnumeratedMs20000
	UeTimersandconstantsT311EnumeratedMs30000
)
