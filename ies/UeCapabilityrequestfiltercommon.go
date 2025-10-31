package ies

import "rrc/utils"

// UE-CapabilityRequestFilterCommon ::= SEQUENCE
// Extensible
type UeCapabilityrequestfiltercommon struct {
	MrdcRequest                 *UeCapabilityrequestfiltercommonMrdcRequest
	CodebooktyperequestR16      *UeCapabilityrequestfiltercommonCodebooktyperequestR16 `ext`
	UplinktxswitchrequestR16    *utils.BOOLEAN                                         `ext`
	RequestedcellgroupingR16    *[]CellgroupingR16                                     `lb:1,ub:maxCellGroupingsR16,ext`
	FallbackgroupfiverequestR17 *utils.BOOLEAN                                         `ext`
}
