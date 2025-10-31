package ies

import "rrc/utils"

// UE-CapabilityRequestFilterNR-v1540 ::= SEQUENCE
type UeCapabilityrequestfilternrV1540 struct {
	SrsSwitchingtimerequest *utils.BOOLEAN
	Noncriticalextension    *UeCapabilityrequestfilternrV1710
}
