package ies

import "rrc/utils"

// SoundingRS-UL-ConfigDedicatedUpPTsExt-r13-setup-transmissionCombNum-r13 ::= ENUMERATED
type SoundingrsUlConfigdedicatedupptsextR13SetupTransmissioncombnumR13 struct {
	Value utils.ENUMERATED
}

const (
	SoundingrsUlConfigdedicatedupptsextR13SetupTransmissioncombnumR13EnumeratedNothing = iota
	SoundingrsUlConfigdedicatedupptsextR13SetupTransmissioncombnumR13EnumeratedN2
	SoundingrsUlConfigdedicatedupptsextR13SetupTransmissioncombnumR13EnumeratedN4
)
