package ies

import "rrc/utils"

// PollPDU-v1310 ::= ENUMERATED
type PollpduV1310 struct {
	Value utils.ENUMERATED
}

const (
	PollpduV1310P512   = 0
	PollpduV1310P1024  = 1
	PollpduV1310P2048  = 2
	PollpduV1310P4096  = 3
	PollpduV1310P6144  = 4
	PollpduV1310P8192  = 5
	PollpduV1310P12288 = 6
	PollpduV1310P16384 = 7
)
