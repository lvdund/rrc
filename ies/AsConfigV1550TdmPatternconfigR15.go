package ies

import "rrc/utils"

// AS-Config-v1550-tdm-PatternConfig-r15 ::= SEQUENCE
type AsConfigV1550TdmPatternconfigR15 struct {
	SubframeassignmentR15 SubframeassignmentR15
	HarqOffsetR15         utils.INTEGER `lb:0,ub:9`
}
