package ies

import "rrc/utils"

// SL-GapPattern-r13 ::= SEQUENCE
// Extensible
type SlGappatternR13 struct {
	GapperiodR13         SlGappatternR13GapperiodR13
	GapoffsetR12         SlOffsetindicatorR12
	GapsubframebitmapR13 utils.BITSTRING `lb:1,ub:10240`
}
