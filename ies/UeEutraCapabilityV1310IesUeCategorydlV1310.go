package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1310-IEs-ue-CategoryDL-v1310 ::= ENUMERATED
type UeEutraCapabilityV1310IesUeCategorydlV1310 struct {
	Value utils.ENUMERATED
}

const (
	UeEutraCapabilityV1310IesUeCategorydlV1310EnumeratedNothing = iota
	UeEutraCapabilityV1310IesUeCategorydlV1310EnumeratedN17
	UeEutraCapabilityV1310IesUeCategorydlV1310EnumeratedM1
)
