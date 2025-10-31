package ies

import "rrc/utils"

// PreambleTransMax ::= ENUMERATED
type Preambletransmax struct {
	Value utils.ENUMERATED
}

const (
	PreambletransmaxEnumeratedNothing = iota
	PreambletransmaxEnumeratedN3
	PreambletransmaxEnumeratedN4
	PreambletransmaxEnumeratedN5
	PreambletransmaxEnumeratedN6
	PreambletransmaxEnumeratedN7
	PreambletransmaxEnumeratedN8
	PreambletransmaxEnumeratedN10
	PreambletransmaxEnumeratedN20
	PreambletransmaxEnumeratedN50
	PreambletransmaxEnumeratedN100
	PreambletransmaxEnumeratedN200
)
