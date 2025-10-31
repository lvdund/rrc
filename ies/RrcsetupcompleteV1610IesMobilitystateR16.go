package ies

import "rrc/utils"

// RRCSetupComplete-v1610-IEs-mobilityState-r16 ::= ENUMERATED
type RrcsetupcompleteV1610IesMobilitystateR16 struct {
	Value utils.ENUMERATED
}

const (
	RrcsetupcompleteV1610IesMobilitystateR16EnumeratedNothing = iota
	RrcsetupcompleteV1610IesMobilitystateR16EnumeratedNormal
	RrcsetupcompleteV1610IesMobilitystateR16EnumeratedMedium
	RrcsetupcompleteV1610IesMobilitystateR16EnumeratedHigh
	RrcsetupcompleteV1610IesMobilitystateR16EnumeratedSpare
)
