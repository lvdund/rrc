package ies

import "rrc/utils"

// SoundingRS-UL-ConfigDedicatedAperiodic-v1430-setup ::= SEQUENCE
type SoundingrsUlConfigdedicatedaperiodicV1430Setup struct {
	SrsSubframeindicationR14 *utils.INTEGER `lb:0,ub:4`
}
