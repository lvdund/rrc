package ies

import "rrc/utils"

// UE-Capability-NB-Ext-r14-IEs-ue-Category-NB-r14 ::= ENUMERATED
type UeCapabilityNbExtR14IesUeCategoryNbR14 struct {
	Value utils.ENUMERATED
}

const (
	UeCapabilityNbExtR14IesUeCategoryNbR14EnumeratedNothing = iota
	UeCapabilityNbExtR14IesUeCategoryNbR14EnumeratedNb2
)
