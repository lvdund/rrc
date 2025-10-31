package ies

import "rrc/utils"

// PollPDU ::= ENUMERATED
type Pollpdu struct {
	Value utils.ENUMERATED
}

const (
	PollpduEnumeratedNothing = iota
	PollpduEnumeratedP4
	PollpduEnumeratedP8
	PollpduEnumeratedP16
	PollpduEnumeratedP32
	PollpduEnumeratedP64
	PollpduEnumeratedP128
	PollpduEnumeratedP256
	PollpduEnumeratedPinfinity
)
