package ies

import "rrc/utils"

// RLC-Parameters-v1530-flexibleUM-AM-Combinations-r15 ::= ENUMERATED
type RlcParametersV1530FlexibleumAmCombinationsR15 struct {
	Value utils.ENUMERATED
}

const (
	RlcParametersV1530FlexibleumAmCombinationsR15EnumeratedNothing = iota
	RlcParametersV1530FlexibleumAmCombinationsR15EnumeratedSupported
)
