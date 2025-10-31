package ies

import "rrc/utils"

// CodebookParameters-v1610-supportedCSI-RS-ResourceListAlt-r16 ::= SEQUENCE
type CodebookparametersV1610SupportedcsiRsResourcelistaltR16 struct {
	Type1SinglepanelR16   *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResources`
	Type1MultipanelR16    *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResources`
	Type2R16              *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResources`
	Type2PortselectionR16 *[]utils.INTEGER `lb:1,ub:maxNrofCSIRsResources`
}
