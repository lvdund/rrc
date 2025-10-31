package ies

import "rrc/utils"

// UE-NR-Capability-v1530-delayBudgetReporting ::= ENUMERATED
type UeNrCapabilityV1530Delaybudgetreporting struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1530DelaybudgetreportingEnumeratedNothing = iota
	UeNrCapabilityV1530DelaybudgetreportingEnumeratedSupported
)
