package ies

import "rrc/utils"

// SoundingRS-UL-ConfigDedicated-setup-srs-Bandwidth ::= ENUMERATED
type SoundingrsUlConfigdedicatedSetupSrsBandwidth struct {
	Value utils.ENUMERATED
}

const (
	SoundingrsUlConfigdedicatedSetupSrsBandwidthEnumeratedNothing = iota
	SoundingrsUlConfigdedicatedSetupSrsBandwidthEnumeratedBw0
	SoundingrsUlConfigdedicatedSetupSrsBandwidthEnumeratedBw1
	SoundingrsUlConfigdedicatedSetupSrsBandwidthEnumeratedBw2
	SoundingrsUlConfigdedicatedSetupSrsBandwidthEnumeratedBw3
)
