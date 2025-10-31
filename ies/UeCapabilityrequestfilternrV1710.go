package ies

import "rrc/utils"

// UE-CapabilityRequestFilterNR-v1710 ::= SEQUENCE
type UeCapabilityrequestfilternrV1710 struct {
	SidelinkrequestR17   *utils.BOOLEAN
	Noncriticalextension *UeCapabilityrequestfilternrV1710Noncriticalextension
}
