package ies

import "rrc/utils"

// SoundingRS-UL-ConfigDedicated-v1310-setup-transmissionCombNum-r13 ::= ENUMERATED
type SoundingrsUlConfigdedicatedV1310SetupTransmissioncombnumR13 struct {
	Value utils.ENUMERATED
}

const (
	SoundingrsUlConfigdedicatedV1310SetupTransmissioncombnumR13EnumeratedNothing = iota
	SoundingrsUlConfigdedicatedV1310SetupTransmissioncombnumR13EnumeratedN2
	SoundingrsUlConfigdedicatedV1310SetupTransmissioncombnumR13EnumeratedN4
)
