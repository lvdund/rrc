package ies

import "rrc/utils"

// PollPDU-r15 ::= ENUMERATED
type PollpduR15 struct {
	Value utils.ENUMERATED
}

const (
	PollpduR15P4        = 0
	PollpduR15P8        = 1
	PollpduR15P16       = 2
	PollpduR15P32       = 3
	PollpduR15P64       = 4
	PollpduR15P128      = 5
	PollpduR15P256      = 6
	PollpduR15P512      = 7
	PollpduR15P1024     = 8
	PollpduR15P2048R15  = 9
	PollpduR15P4096R15  = 10
	PollpduR15P6144R15  = 11
	PollpduR15P8192R15  = 12
	PollpduR15P12288R15 = 13
	PollpduR15P16384R15 = 14
	PollpduR15Pinfinity = 15
)
