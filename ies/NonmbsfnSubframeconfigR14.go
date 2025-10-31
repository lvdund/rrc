package ies

import "rrc/utils"

// NonMBSFN-SubframeConfig-r14 ::= SEQUENCE
type NonmbsfnSubframeconfigR14 struct {
	RadioframeallocationperiodR14 NonmbsfnSubframeconfigR14RadioframeallocationperiodR14
	RadioframeallocationoffsetR14 utils.INTEGER   `lb:0,ub:7`
	SubframeallocationR14         utils.BITSTRING `lb:9,ub:9`
}
