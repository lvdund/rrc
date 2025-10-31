package ies

import "rrc/utils"

// DownlinkPreemption-timeFrequencySet ::= ENUMERATED
type DownlinkpreemptionTimefrequencyset struct {
	Value utils.ENUMERATED
}

const (
	DownlinkpreemptionTimefrequencysetEnumeratedNothing = iota
	DownlinkpreemptionTimefrequencysetEnumeratedSet0
	DownlinkpreemptionTimefrequencysetEnumeratedSet1
)
