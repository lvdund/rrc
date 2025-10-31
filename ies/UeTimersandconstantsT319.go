package ies

import "rrc/utils"

// UE-TimersAndConstants-t319 ::= ENUMERATED
type UeTimersandconstantsT319 struct {
	Value utils.ENUMERATED
}

const (
	UeTimersandconstantsT319EnumeratedNothing = iota
	UeTimersandconstantsT319EnumeratedMs100
	UeTimersandconstantsT319EnumeratedMs200
	UeTimersandconstantsT319EnumeratedMs300
	UeTimersandconstantsT319EnumeratedMs400
	UeTimersandconstantsT319EnumeratedMs600
	UeTimersandconstantsT319EnumeratedMs1000
	UeTimersandconstantsT319EnumeratedMs1500
	UeTimersandconstantsT319EnumeratedMs2000
)
