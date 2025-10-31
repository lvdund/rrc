package ies

import "rrc/utils"

// RRCConnectionSetupComplete-NB-v1430-IEs-gummei-Type-r14 ::= ENUMERATED
type RrcconnectionsetupcompleteNbV1430IesGummeiTypeR14 struct {
	Value utils.ENUMERATED
}

const (
	RrcconnectionsetupcompleteNbV1430IesGummeiTypeR14EnumeratedNothing = iota
	RrcconnectionsetupcompleteNbV1430IesGummeiTypeR14EnumeratedMapped
)
