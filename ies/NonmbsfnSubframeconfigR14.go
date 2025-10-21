package ies

import "rrc/utils"

// NonMBSFN-SubframeConfig-r14 ::= SEQUENCE
type NonmbsfnSubframeconfigR14 struct {
	RadioframeallocationperiodR14 utils.ENUMERATED
	RadioframeallocationoffsetR14 utils.INTEGER
	SubframeallocationR14         utils.BITSTRING
}
