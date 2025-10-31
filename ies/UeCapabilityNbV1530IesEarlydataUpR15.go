package ies

import "rrc/utils"

// UE-Capability-NB-v1530-IEs-earlyData-UP-r15 ::= ENUMERATED
type UeCapabilityNbV1530IesEarlydataUpR15 struct {
	Value utils.ENUMERATED
}

const (
	UeCapabilityNbV1530IesEarlydataUpR15EnumeratedNothing = iota
	UeCapabilityNbV1530IesEarlydataUpR15EnumeratedSupported
)
