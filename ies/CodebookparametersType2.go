package ies

import "rrc/utils"

// CodebookParameters-type2 ::= SEQUENCE
type CodebookparametersType2 struct {
	SupportedcsiRsResourcelist []SupportedcsiRsResource `lb:1,ub:maxNrofCSIRsResources`
	Parameterlx                utils.INTEGER            `lb:0,ub:4`
	Amplitudescalingtype       CodebookparametersType2Amplitudescalingtype
	Amplitudesubsetrestriction *CodebookparametersType2Amplitudesubsetrestriction
}
