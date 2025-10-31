package ies

import "rrc/utils"

// RLF-TimersAndConstants-t311 ::= ENUMERATED
type RlfTimersandconstantsT311 struct {
	Value utils.ENUMERATED
}

const (
	RlfTimersandconstantsT311EnumeratedNothing = iota
	RlfTimersandconstantsT311EnumeratedMs1000
	RlfTimersandconstantsT311EnumeratedMs3000
	RlfTimersandconstantsT311EnumeratedMs5000
	RlfTimersandconstantsT311EnumeratedMs10000
	RlfTimersandconstantsT311EnumeratedMs15000
	RlfTimersandconstantsT311EnumeratedMs20000
	RlfTimersandconstantsT311EnumeratedMs30000
)
