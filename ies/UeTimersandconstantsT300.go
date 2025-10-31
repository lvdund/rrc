package ies

import "rrc/utils"

// UE-TimersAndConstants-t300 ::= ENUMERATED
type UeTimersandconstantsT300 struct {
	Value utils.ENUMERATED
}

const (
	UeTimersandconstantsT300EnumeratedNothing = iota
	UeTimersandconstantsT300EnumeratedMs100
	UeTimersandconstantsT300EnumeratedMs200
	UeTimersandconstantsT300EnumeratedMs300
	UeTimersandconstantsT300EnumeratedMs400
	UeTimersandconstantsT300EnumeratedMs600
	UeTimersandconstantsT300EnumeratedMs1000
	UeTimersandconstantsT300EnumeratedMs1500
	UeTimersandconstantsT300EnumeratedMs2000
)
