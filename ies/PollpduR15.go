package ies

import "rrc/utils"

// PollPDU-r15 ::= ENUMERATED
type PollpduR15 struct {
	Value utils.ENUMERATED
}

const (
	PollpduR15EnumeratedNothing = iota
	PollpduR15EnumeratedP4
	PollpduR15EnumeratedP8
	PollpduR15EnumeratedP16
	PollpduR15EnumeratedP32
	PollpduR15EnumeratedP64
	PollpduR15EnumeratedP128
	PollpduR15EnumeratedP256
	PollpduR15EnumeratedP512
	PollpduR15EnumeratedP1024
	PollpduR15EnumeratedP2048_R15
	PollpduR15EnumeratedP4096_R15
	PollpduR15EnumeratedP6144_R15
	PollpduR15EnumeratedP8192_R15
	PollpduR15EnumeratedP12288_R15
	PollpduR15EnumeratedP16384_R15
	PollpduR15EnumeratedPinfinity
)
