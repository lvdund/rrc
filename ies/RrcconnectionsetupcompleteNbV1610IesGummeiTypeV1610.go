package ies

import "rrc/utils"

// RRCConnectionSetupComplete-NB-v1610-IEs-gummei-Type-v1610 ::= ENUMERATED
type RrcconnectionsetupcompleteNbV1610IesGummeiTypeV1610 struct {
	Value utils.ENUMERATED
}

const (
	RrcconnectionsetupcompleteNbV1610IesGummeiTypeV1610EnumeratedNothing = iota
	RrcconnectionsetupcompleteNbV1610IesGummeiTypeV1610EnumeratedMappedfrom5g
)
