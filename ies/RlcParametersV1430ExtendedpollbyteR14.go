package ies

import "rrc/utils"

// RLC-Parameters-v1430-extendedPollByte-r14 ::= ENUMERATED
type RlcParametersV1430ExtendedpollbyteR14 struct {
	Value utils.ENUMERATED
}

const (
	RlcParametersV1430ExtendedpollbyteR14EnumeratedNothing = iota
	RlcParametersV1430ExtendedpollbyteR14EnumeratedSupported
)
