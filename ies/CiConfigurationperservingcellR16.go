package ies

import "rrc/utils"

// CI-ConfigurationPerServingCell-r16 ::= SEQUENCE
// Extensible
type CiConfigurationperservingcellR16 struct {
	Servingcellid                   Servcellindex
	PositionindciR16                utils.INTEGER  `lb:0,ub:maxCIDciPayloadsize1R16`
	PositionindciForsulR16          *utils.INTEGER `lb:0,ub:maxCIDciPayloadsize1R16`
	CiPayloadsizeR16                CiConfigurationperservingcellR16CiPayloadsizeR16
	TimefrequencyregionR16          *CiConfigurationperservingcellR16TimefrequencyregionR16
	UplinkcancellationpriorityV1610 *CiConfigurationperservingcellR16UplinkcancellationpriorityV1610
}
