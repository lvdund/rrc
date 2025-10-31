package ies

import "rrc/utils"

// SoundingRS-UL-ConfigDedicated-setup ::= SEQUENCE
type SoundingrsUlConfigdedicatedSetup struct {
	SrsBandwidth        SoundingrsUlConfigdedicatedSetupSrsBandwidth
	SrsHoppingbandwidth SoundingrsUlConfigdedicatedSetupSrsHoppingbandwidth
	Freqdomainposition  utils.INTEGER `lb:0,ub:23`
	Duration            utils.BOOLEAN
	SrsConfigindex      utils.INTEGER `lb:0,ub:1023`
	Transmissioncomb    utils.INTEGER `lb:0,ub:1`
	Cyclicshift         SoundingrsUlConfigdedicatedSetupCyclicshift
}
