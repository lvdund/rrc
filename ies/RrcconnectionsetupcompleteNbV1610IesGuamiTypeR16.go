package ies

import "rrc/utils"

// RRCConnectionSetupComplete-NB-v1610-IEs-guami-Type-r16 ::= ENUMERATED
type RrcconnectionsetupcompleteNbV1610IesGuamiTypeR16 struct {
	Value utils.ENUMERATED
}

const (
	RrcconnectionsetupcompleteNbV1610IesGuamiTypeR16EnumeratedNothing = iota
	RrcconnectionsetupcompleteNbV1610IesGuamiTypeR16EnumeratedNative
	RrcconnectionsetupcompleteNbV1610IesGuamiTypeR16EnumeratedMapped
)
