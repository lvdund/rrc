package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1350-IEs-ue-CategoryUL-v1350 ::= ENUMERATED
type UeEutraCapabilityV1350IesUeCategoryulV1350 struct {
	Value utils.ENUMERATED
}

const (
	UeEutraCapabilityV1350IesUeCategoryulV1350EnumeratedNothing = iota
	UeEutraCapabilityV1350IesUeCategoryulV1350EnumeratedOnebis
)
