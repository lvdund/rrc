package ies

import "rrc/utils"

// CSI-RS-ResourceMapping ::= SEQUENCE
// Extensible
type CsiRsResourcemapping struct {
	Frequencydomainallocation    CsiRsResourcemappingFrequencydomainallocation
	Nrofports                    CsiRsResourcemappingNrofports
	Firstofdmsymbolintimedomain  utils.INTEGER  `lb:0,ub:13`
	Firstofdmsymbolintimedomain2 *utils.INTEGER `lb:0,ub:12`
	CdmType                      CsiRsResourcemappingCdmType
	Density                      CsiRsResourcemappingDensity
	Freqband                     CsiFrequencyoccupation
}
