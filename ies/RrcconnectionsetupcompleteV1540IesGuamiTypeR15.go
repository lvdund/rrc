package ies

import "rrc/utils"

// RRCConnectionSetupComplete-v1540-IEs-guami-Type-r15 ::= ENUMERATED
type RrcconnectionsetupcompleteV1540IesGuamiTypeR15 struct {
	Value utils.ENUMERATED
}

const (
	RrcconnectionsetupcompleteV1540IesGuamiTypeR15EnumeratedNothing = iota
	RrcconnectionsetupcompleteV1540IesGuamiTypeR15EnumeratedNative
	RrcconnectionsetupcompleteV1540IesGuamiTypeR15EnumeratedMapped
)
