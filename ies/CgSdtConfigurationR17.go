package ies

import "rrc/utils"

// CG-SDT-Configuration-r17 ::= SEQUENCE
type CgSdtConfigurationR17 struct {
	CgSdtRetransmissiontimer *utils.INTEGER `lb:0,ub:64`
	SdtSsbSubsetR17          *CgSdtConfigurationR17SdtSsbSubsetR17
	SdtSsbPercgPuschR17      *CgSdtConfigurationR17SdtSsbPercgPuschR17
	SdtP0PuschR17            *utils.INTEGER `lb:0,ub:15`
	SdtAlphaR17              *CgSdtConfigurationR17SdtAlphaR17
	SdtDmrsPortsR17          *CgSdtConfigurationR17SdtDmrsPortsR17
	SdtNrofdmrsSequencesR17  *utils.INTEGER `lb:0,ub:2`
}
