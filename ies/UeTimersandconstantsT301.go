package ies

import "rrc/utils"

// UE-TimersAndConstants-t301 ::= ENUMERATED
type UeTimersandconstantsT301 struct {
	Value utils.ENUMERATED
}

const (
	UeTimersandconstantsT301EnumeratedNothing = iota
	UeTimersandconstantsT301EnumeratedMs100
	UeTimersandconstantsT301EnumeratedMs200
	UeTimersandconstantsT301EnumeratedMs300
	UeTimersandconstantsT301EnumeratedMs400
	UeTimersandconstantsT301EnumeratedMs600
	UeTimersandconstantsT301EnumeratedMs1000
	UeTimersandconstantsT301EnumeratedMs1500
	UeTimersandconstantsT301EnumeratedMs2000
)
