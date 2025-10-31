package ies

import "rrc/utils"

// TDM-PatternConfig-r15-setup ::= SEQUENCE
type TdmPatternconfigR15Setup struct {
	SubframeassignmentR15 SubframeassignmentR15
	HarqOffsetR15         utils.INTEGER `lb:0,ub:9`
}
