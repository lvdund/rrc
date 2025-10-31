package ies

import "rrc/utils"

// RRCConnectionResume-r13-IEs ::= SEQUENCE
type RrcconnectionresumeR13 struct {
	RadioresourceconfigdedicatedR13 *Radioresourceconfigdedicated
	NexthopchainingcountR13         Nexthopchainingcount
	MeasconfigR13                   *Measconfig
	AntennainfodedicatedpcellR13    *AntennainfodedicatedV10i0
	DrbContinuerohcR13              *bool
	Latenoncriticalextension        *utils.OCTETSTRING
	RrcconnectionresumeV1430        *RrcconnectionresumeV1430
}
