package ies

import "rrc/utils"

// PreambleTransMax ::= ENUMERATED
type Preambletransmax struct {
	Value utils.ENUMERATED
}

const (
	PreambletransmaxN3   = 0
	PreambletransmaxN4   = 1
	PreambletransmaxN5   = 2
	PreambletransmaxN6   = 3
	PreambletransmaxN7   = 4
	PreambletransmaxN8   = 5
	PreambletransmaxN10  = 6
	PreambletransmaxN20  = 7
	PreambletransmaxN50  = 8
	PreambletransmaxN100 = 9
	PreambletransmaxN200 = 10
)
