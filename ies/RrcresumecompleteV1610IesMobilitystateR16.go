package ies

import "rrc/utils"

// RRCResumeComplete-v1610-IEs-mobilityState-r16 ::= ENUMERATED
type RrcresumecompleteV1610IesMobilitystateR16 struct {
	Value utils.ENUMERATED
}

const (
	RrcresumecompleteV1610IesMobilitystateR16EnumeratedNothing = iota
	RrcresumecompleteV1610IesMobilitystateR16EnumeratedNormal
	RrcresumecompleteV1610IesMobilitystateR16EnumeratedMedium
	RrcresumecompleteV1610IesMobilitystateR16EnumeratedHigh
	RrcresumecompleteV1610IesMobilitystateR16EnumeratedSpare
)
