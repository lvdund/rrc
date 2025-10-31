package ies

import "rrc/utils"

// P-a ::= ENUMERATED
type PA struct {
	Value utils.ENUMERATED
}

const (
	PAEnumeratedNothing = iota
	PAEnumeratedDb_6
	PAEnumeratedDb_4dot77
	PAEnumeratedDb_3
	PAEnumeratedDb_1dot77
	PAEnumeratedDb0
	PAEnumeratedDb1
	PAEnumeratedDb2
	PAEnumeratedDb3
)
