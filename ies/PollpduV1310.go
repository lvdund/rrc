package ies

import "rrc/utils"

// PollPDU-v1310 ::= ENUMERATED
type PollpduV1310 struct {
	Value utils.ENUMERATED
}

const (
	PollpduV1310EnumeratedNothing = iota
	PollpduV1310EnumeratedP512
	PollpduV1310EnumeratedP1024
	PollpduV1310EnumeratedP2048
	PollpduV1310EnumeratedP4096
	PollpduV1310EnumeratedP6144
	PollpduV1310EnumeratedP8192
	PollpduV1310EnumeratedP12288
	PollpduV1310EnumeratedP16384
)
