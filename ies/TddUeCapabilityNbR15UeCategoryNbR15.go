package ies

import "rrc/utils"

// TDD-UE-Capability-NB-r15-ue-Category-NB-r15 ::= ENUMERATED
type TddUeCapabilityNbR15UeCategoryNbR15 struct {
	Value utils.ENUMERATED
}

const (
	TddUeCapabilityNbR15UeCategoryNbR15EnumeratedNothing = iota
	TddUeCapabilityNbR15UeCategoryNbR15EnumeratedNb2
)
