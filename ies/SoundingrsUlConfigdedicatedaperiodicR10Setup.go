package ies

import "rrc/utils"

// SoundingRS-UL-ConfigDedicatedAperiodic-r10-setup ::= SEQUENCE
// Extensible
type SoundingrsUlConfigdedicatedaperiodicR10Setup struct {
	SrsConfigindexapR10      utils.INTEGER     `lb:0,ub:31`
	SrsConfigapdciFormat4R10 *[]SrsConfigapR10 `lb:1,ub:3`
	SrsActivateapR10         *SoundingrsUlConfigdedicatedaperiodicR10SetupSrsActivateapR10
}
