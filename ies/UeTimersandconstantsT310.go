package ies

import "rrc/utils"

// UE-TimersAndConstants-t310 ::= ENUMERATED
type UeTimersandconstantsT310 struct {
	Value utils.ENUMERATED
}

const (
	UeTimersandconstantsT310EnumeratedNothing = iota
	UeTimersandconstantsT310EnumeratedMs0
	UeTimersandconstantsT310EnumeratedMs50
	UeTimersandconstantsT310EnumeratedMs100
	UeTimersandconstantsT310EnumeratedMs200
	UeTimersandconstantsT310EnumeratedMs500
	UeTimersandconstantsT310EnumeratedMs1000
	UeTimersandconstantsT310EnumeratedMs2000
)
