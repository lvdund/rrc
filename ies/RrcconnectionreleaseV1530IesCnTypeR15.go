package ies

import "rrc/utils"

// RRCConnectionRelease-v1530-IEs-cn-Type-r15 ::= ENUMERATED
type RrcconnectionreleaseV1530IesCnTypeR15 struct {
	Value utils.ENUMERATED
}

const (
	RrcconnectionreleaseV1530IesCnTypeR15EnumeratedNothing = iota
	RrcconnectionreleaseV1530IesCnTypeR15EnumeratedEpc
	RrcconnectionreleaseV1530IesCnTypeR15EnumeratedFivegc
)
