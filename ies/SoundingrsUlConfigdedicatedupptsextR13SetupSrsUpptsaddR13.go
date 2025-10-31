package ies

import "rrc/utils"

// SoundingRS-UL-ConfigDedicatedUpPTsExt-r13-setup-srs-UpPtsAdd-r13 ::= ENUMERATED
type SoundingrsUlConfigdedicatedupptsextR13SetupSrsUpptsaddR13 struct {
	Value utils.ENUMERATED
}

const (
	SoundingrsUlConfigdedicatedupptsextR13SetupSrsUpptsaddR13EnumeratedNothing = iota
	SoundingrsUlConfigdedicatedupptsextR13SetupSrsUpptsaddR13EnumeratedSym2
	SoundingrsUlConfigdedicatedupptsextR13SetupSrsUpptsaddR13EnumeratedSym4
)
