package ies

import "rrc/utils"

// CFI-PatternConfig-r15 ::= SEQUENCE
type CfiPatternconfigR15 struct {
	CfiPatternsubframeR15    *[]utils.INTEGER `lb:10,ub:10`
	CfiPatternslotsubslotR15 *[]utils.INTEGER `lb:10,ub:10`
}
