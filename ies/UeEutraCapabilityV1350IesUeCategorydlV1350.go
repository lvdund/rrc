package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1350-IEs-ue-CategoryDL-v1350 ::= ENUMERATED
type UeEutraCapabilityV1350IesUeCategorydlV1350 struct {
	Value utils.ENUMERATED
}

const (
	UeEutraCapabilityV1350IesUeCategorydlV1350EnumeratedNothing = iota
	UeEutraCapabilityV1350IesUeCategorydlV1350EnumeratedOnebis
)
