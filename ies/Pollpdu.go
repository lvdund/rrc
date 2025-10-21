package ies

import "rrc/utils"

// PollPDU ::= ENUMERATED
type Pollpdu struct {
	Value utils.ENUMERATED
}

const (
	PollpduP4        = 0
	PollpduP8        = 1
	PollpduP16       = 2
	PollpduP32       = 3
	PollpduP64       = 4
	PollpduP128      = 5
	PollpduP256      = 6
	PollpduPinfinity = 7
)
