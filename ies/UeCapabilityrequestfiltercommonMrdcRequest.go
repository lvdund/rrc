package ies

import "rrc/utils"

// UE-CapabilityRequestFilterCommon-mrdc-Request ::= SEQUENCE
type UeCapabilityrequestfiltercommonMrdcRequest struct {
	OmitenDc    *utils.BOOLEAN
	IncludenrDc *utils.BOOLEAN
	IncludeneDc *utils.BOOLEAN
}
