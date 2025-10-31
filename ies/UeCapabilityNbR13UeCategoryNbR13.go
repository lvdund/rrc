package ies

import "rrc/utils"

// UE-Capability-NB-r13-ue-Category-NB-r13 ::= ENUMERATED
type UeCapabilityNbR13UeCategoryNbR13 struct {
	Value utils.ENUMERATED
}

const (
	UeCapabilityNbR13UeCategoryNbR13EnumeratedNothing = iota
	UeCapabilityNbR13UeCategoryNbR13EnumeratedNb1
)
