package ies

import "rrc/utils"

// RRCConnectionSetupComplete-v1320-IEs-ce-ModeB-r13 ::= ENUMERATED
type RrcconnectionsetupcompleteV1320IesCeModebR13 struct {
	Value utils.ENUMERATED
}

const (
	RrcconnectionsetupcompleteV1320IesCeModebR13EnumeratedNothing = iota
	RrcconnectionsetupcompleteV1320IesCeModebR13EnumeratedSupported
)
