package ies

import "rrc/utils"

// SoundingRS-UL-ConfigDedicatedAperiodicUpPTsExt-r13-setup ::= SEQUENCE
type SoundingrsUlConfigdedicatedaperiodicupptsextR13Setup struct {
	SrsUpptsaddR13           SoundingrsUlConfigdedicatedaperiodicupptsextR13SetupSrsUpptsaddR13
	SrsConfigindexapR13      utils.INTEGER     `lb:0,ub:31`
	SrsConfigapdciFormat4R13 *[]SrsConfigapR13 `lb:1,ub:3`
	SrsActivateapR13         *SoundingrsUlConfigdedicatedaperiodicupptsextR13SetupSrsActivateapR13
}
