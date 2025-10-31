package ies

import "rrc/utils"

// UE-NR-Capability-v1610-resumeWithStoredMCG-SCells-r16 ::= ENUMERATED
type UeNrCapabilityV1610ResumewithstoredmcgScellsR16 struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1610ResumewithstoredmcgScellsR16EnumeratedNothing = iota
	UeNrCapabilityV1610ResumewithstoredmcgScellsR16EnumeratedSupported
)
