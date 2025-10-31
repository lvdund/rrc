package ies

import "rrc/utils"

// CodebookParameters-type2-PortSelection ::= SEQUENCE
type CodebookparametersType2Portselection struct {
	SupportedcsiRsResourcelist []SupportedcsiRsResource `lb:1,ub:maxNrofCSIRsResources`
	Parameterlx                utils.INTEGER            `lb:0,ub:4`
	Amplitudescalingtype       CodebookparametersType2PortselectionAmplitudescalingtype
}
