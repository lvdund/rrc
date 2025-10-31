package ies

import "rrc/utils"

// CG-StartingOffsets-r16 ::= SEQUENCE
type CgStartingoffsetsR16 struct {
	CgStartingfullbwInsidecotR16     *[]utils.INTEGER `lb:1,ub:7`
	CgStartingfullbwOutsidecotR16    *[]utils.INTEGER `lb:1,ub:7`
	CgStartingpartialbwInsidecotR16  *utils.INTEGER   `lb:0,ub:6`
	CgStartingpartialbwOutsidecotR16 *utils.INTEGER   `lb:0,ub:6`
}
