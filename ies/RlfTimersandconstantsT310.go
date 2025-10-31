package ies

import "rrc/utils"

// RLF-TimersAndConstants-t310 ::= ENUMERATED
type RlfTimersandconstantsT310 struct {
	Value utils.ENUMERATED
}

const (
	RlfTimersandconstantsT310EnumeratedNothing = iota
	RlfTimersandconstantsT310EnumeratedMs0
	RlfTimersandconstantsT310EnumeratedMs50
	RlfTimersandconstantsT310EnumeratedMs100
	RlfTimersandconstantsT310EnumeratedMs200
	RlfTimersandconstantsT310EnumeratedMs500
	RlfTimersandconstantsT310EnumeratedMs1000
	RlfTimersandconstantsT310EnumeratedMs2000
	RlfTimersandconstantsT310EnumeratedMs4000
	RlfTimersandconstantsT310EnumeratedMs6000
)
