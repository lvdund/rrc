package ies

import "rrc/utils"

// SoundingRS-UL-ConfigDedicatedUpPTsExt-r13-setup ::= SEQUENCE
type SoundingrsUlConfigdedicatedupptsextR13Setup struct {
	SrsUpptsaddR13         SoundingrsUlConfigdedicatedupptsextR13SetupSrsUpptsaddR13
	SrsBandwidthR13        SoundingrsUlConfigdedicatedupptsextR13SetupSrsBandwidthR13
	SrsHoppingbandwidthR13 SoundingrsUlConfigdedicatedupptsextR13SetupSrsHoppingbandwidthR13
	FreqdomainpositionR13  utils.INTEGER `lb:0,ub:23`
	DurationR13            utils.BOOLEAN
	SrsConfigindexR13      utils.INTEGER `lb:0,ub:1023`
	TransmissioncombR13    utils.INTEGER `lb:0,ub:3`
	CyclicshiftR13         SoundingrsUlConfigdedicatedupptsextR13SetupCyclicshiftR13
	SrsAntennaportR13      SrsAntennaport
	TransmissioncombnumR13 SoundingrsUlConfigdedicatedupptsextR13SetupTransmissioncombnumR13
}
