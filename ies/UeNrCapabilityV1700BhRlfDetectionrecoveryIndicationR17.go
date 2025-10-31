package ies

import "rrc/utils"

// UE-NR-Capability-v1700-bh-RLF-DetectionRecovery-Indication-r17 ::= ENUMERATED
type UeNrCapabilityV1700BhRlfDetectionrecoveryIndicationR17 struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1700BhRlfDetectionrecoveryIndicationR17EnumeratedNothing = iota
	UeNrCapabilityV1700BhRlfDetectionrecoveryIndicationR17EnumeratedSupported
)
