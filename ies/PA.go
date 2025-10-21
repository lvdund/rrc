package ies

import "rrc/utils"

// P-a ::= ENUMERATED
type PA struct {
	Value utils.ENUMERATED
}

const (
	PADb6      = 0
	PADb4dot77 = 1
	PADb3      = 2
	PADb1dot77 = 3
	PADb0      = 4
	PADb1      = 5
	PADb2      = 6
	PADb3      = 7
)
