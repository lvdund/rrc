package ies

import "rrc/utils"

// CodebookParameters-type1-singlePanel ::= SEQUENCE
type CodebookparametersType1Singlepanel struct {
	SupportedcsiRsResourcelist   []SupportedcsiRsResource `lb:1,ub:maxNrofCSIRsResources`
	Modes                        CodebookparametersType1SinglepanelModes
	MaxnumbercsiRsPerresourceset utils.INTEGER `lb:0,ub:8`
}
