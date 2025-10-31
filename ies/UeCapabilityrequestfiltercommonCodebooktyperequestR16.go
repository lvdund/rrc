package ies

import "rrc/utils"

// UE-CapabilityRequestFilterCommon-codebookTypeRequest-r16 ::= SEQUENCE
type UeCapabilityrequestfiltercommonCodebooktyperequestR16 struct {
	Type1SinglepanelR16   *utils.BOOLEAN
	Type1MultipanelR16    *utils.BOOLEAN
	Type2R16              *utils.BOOLEAN
	Type2PortselectionR16 *utils.BOOLEAN
}
