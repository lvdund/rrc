package ies

import "rrc/utils"

// RRCConnectionSetupComplete-v1020-IEs-gummei-Type-r10 ::= ENUMERATED
type RrcconnectionsetupcompleteV1020IesGummeiTypeR10 struct {
	Value utils.ENUMERATED
}

const (
	RrcconnectionsetupcompleteV1020IesGummeiTypeR10EnumeratedNothing = iota
	RrcconnectionsetupcompleteV1020IesGummeiTypeR10EnumeratedNative
	RrcconnectionsetupcompleteV1020IesGummeiTypeR10EnumeratedMapped
)
