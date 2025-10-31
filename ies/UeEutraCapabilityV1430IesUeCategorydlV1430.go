package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1430-IEs-ue-CategoryDL-v1430 ::= ENUMERATED
type UeEutraCapabilityV1430IesUeCategorydlV1430 struct {
	Value utils.ENUMERATED
}

const (
	UeEutraCapabilityV1430IesUeCategorydlV1430EnumeratedNothing = iota
	UeEutraCapabilityV1430IesUeCategorydlV1430EnumeratedM2
)
