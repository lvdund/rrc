package ies

import "rrc/utils"

// CodebookParameters-type1-multiPanel ::= SEQUENCE
type CodebookparametersType1Multipanel struct {
	SupportedcsiRsResourcelist   []SupportedcsiRsResource `lb:1,ub:maxNrofCSIRsResources`
	Modes                        CodebookparametersType1MultipanelModes
	Nrofpanels                   CodebookparametersType1MultipanelNrofpanels
	MaxnumbercsiRsPerresourceset utils.INTEGER `lb:0,ub:8`
}
