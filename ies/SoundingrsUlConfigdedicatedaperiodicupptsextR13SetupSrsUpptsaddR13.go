package ies

import "rrc/utils"

// SoundingRS-UL-ConfigDedicatedAperiodicUpPTsExt-r13-setup-srs-UpPtsAdd-r13 ::= ENUMERATED
type SoundingrsUlConfigdedicatedaperiodicupptsextR13SetupSrsUpptsaddR13 struct {
	Value utils.ENUMERATED
}

const (
	SoundingrsUlConfigdedicatedaperiodicupptsextR13SetupSrsUpptsaddR13EnumeratedNothing = iota
	SoundingrsUlConfigdedicatedaperiodicupptsextR13SetupSrsUpptsaddR13EnumeratedSym2
	SoundingrsUlConfigdedicatedaperiodicupptsextR13SetupSrsUpptsaddR13EnumeratedSym4
)
