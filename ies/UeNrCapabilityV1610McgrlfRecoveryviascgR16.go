package ies

import "rrc/utils"

// UE-NR-Capability-v1610-mcgRLF-RecoveryViaSCG-r16 ::= ENUMERATED
type UeNrCapabilityV1610McgrlfRecoveryviascgR16 struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1610McgrlfRecoveryviascgR16EnumeratedNothing = iota
	UeNrCapabilityV1610McgrlfRecoveryviascgR16EnumeratedSupported
)
