package ies

import "rrc/utils"

// UE-NR-Capability-v1640-redirectAtResumeByNAS-r16 ::= ENUMERATED
type UeNrCapabilityV1640RedirectatresumebynasR16 struct {
	Value utils.ENUMERATED
}

const (
	UeNrCapabilityV1640RedirectatresumebynasR16EnumeratedNothing = iota
	UeNrCapabilityV1640RedirectatresumebynasR16EnumeratedSupported
)
