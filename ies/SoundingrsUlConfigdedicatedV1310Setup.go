package ies

import "rrc/utils"

// SoundingRS-UL-ConfigDedicated-v1310-setup ::= SEQUENCE
type SoundingrsUlConfigdedicatedV1310Setup struct {
	TransmissioncombV1310  *utils.INTEGER `lb:0,ub:3`
	CyclicshiftV1310       *SoundingrsUlConfigdedicatedV1310SetupCyclicshiftV1310
	TransmissioncombnumR13 *SoundingrsUlConfigdedicatedV1310SetupTransmissioncombnumR13
}
