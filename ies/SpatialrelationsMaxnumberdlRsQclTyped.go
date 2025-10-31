package ies

import "rrc/utils"

// SpatialRelations-maxNumberDL-RS-QCL-TypeD ::= ENUMERATED
type SpatialrelationsMaxnumberdlRsQclTyped struct {
	Value utils.ENUMERATED
}

const (
	SpatialrelationsMaxnumberdlRsQclTypedEnumeratedNothing = iota
	SpatialrelationsMaxnumberdlRsQclTypedEnumeratedN1
	SpatialrelationsMaxnumberdlRsQclTypedEnumeratedN2
	SpatialrelationsMaxnumberdlRsQclTypedEnumeratedN4
	SpatialrelationsMaxnumberdlRsQclTypedEnumeratedN8
	SpatialrelationsMaxnumberdlRsQclTypedEnumeratedN14
)
