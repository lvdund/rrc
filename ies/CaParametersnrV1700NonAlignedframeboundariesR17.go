package ies

import "rrc/utils"

// CA-ParametersNR-v1700-non-AlignedFrameBoundaries-r17 ::= SEQUENCE
type CaParametersnrV1700NonAlignedframeboundariesR17 struct {
	Scs15khz15khzR17 *utils.BITSTRING `lb:1,ub:496`
	Scs15khz30khzR17 *utils.BITSTRING `lb:1,ub:496`
	Scs15khz60khzR17 *utils.BITSTRING `lb:1,ub:496`
	Scs30khz30khzR17 *utils.BITSTRING `lb:1,ub:496`
	Scs30khz60khzR17 *utils.BITSTRING `lb:1,ub:496`
	Scs60khz60khzR17 *utils.BITSTRING `lb:1,ub:496`
}
