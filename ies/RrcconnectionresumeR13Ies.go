package ies

import "rrc/utils"

// RRCConnectionResume-r13-IEs ::= SEQUENCE
type RrcconnectionresumeR13Ies struct {
	RadioresourceconfigdedicatedR13 *Radioresourceconfigdedicated
	NexthopchainingcountR13         Nexthopchainingcount
	MeasconfigR13                   *Measconfig
	AntennainfodedicatedpcellR13    *AntennainfodedicatedV10i0
	DrbContinuerohcR13              *utils.ENUMERATED
	Latenoncriticalextension        *utils.OCTETSTRING
	RrcconnectionresumeV1430Ies     *RrcconnectionresumeV1430Ies
}
