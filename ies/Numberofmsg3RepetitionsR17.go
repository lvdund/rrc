package ies

import "rrc/utils"

// NumberOfMsg3-Repetitions-r17 ::= ENUMERATED
type Numberofmsg3RepetitionsR17 struct {
	Value utils.ENUMERATED
}

const (
	Numberofmsg3RepetitionsR17EnumeratedNothing = iota
	Numberofmsg3RepetitionsR17EnumeratedN1
	Numberofmsg3RepetitionsR17EnumeratedN2
	Numberofmsg3RepetitionsR17EnumeratedN3
	Numberofmsg3RepetitionsR17EnumeratedN4
	Numberofmsg3RepetitionsR17EnumeratedN7
	Numberofmsg3RepetitionsR17EnumeratedN8
	Numberofmsg3RepetitionsR17EnumeratedN12
	Numberofmsg3RepetitionsR17EnumeratedN16
)
