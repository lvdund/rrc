package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1310-IEs-ue-CategoryUL-v1310 ::= ENUMERATED
type UeEutraCapabilityV1310IesUeCategoryulV1310 struct {
	Value utils.ENUMERATED
}

const (
	UeEutraCapabilityV1310IesUeCategoryulV1310EnumeratedNothing = iota
	UeEutraCapabilityV1310IesUeCategoryulV1310EnumeratedN14
	UeEutraCapabilityV1310IesUeCategoryulV1310EnumeratedM1
)
