package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1430-IEs-ue-CategoryUL-v1430b ::= ENUMERATED
type UeEutraCapabilityV1430IesUeCategoryulV1430b struct {
	Value utils.ENUMERATED
}

const (
	UeEutraCapabilityV1430IesUeCategoryulV1430bEnumeratedNothing = iota
	UeEutraCapabilityV1430IesUeCategoryulV1430bEnumeratedN21
)
