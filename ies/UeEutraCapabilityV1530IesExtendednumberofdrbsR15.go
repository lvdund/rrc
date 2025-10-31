package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1530-IEs-extendedNumberOfDRBs-r15 ::= ENUMERATED
type UeEutraCapabilityV1530IesExtendednumberofdrbsR15 struct {
	Value utils.ENUMERATED
}

const (
	UeEutraCapabilityV1530IesExtendednumberofdrbsR15EnumeratedNothing = iota
	UeEutraCapabilityV1530IesExtendednumberofdrbsR15EnumeratedSupported
)
