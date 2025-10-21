package ies

import "rrc/utils"

// MeasResultSSB-IndexList-r15 ::= SEQUENCE OF MeasResultSSB-Index-r15
// SIZE (1..maxRS-IndexReport-r15)
type MeasresultssbIndexlistR15 struct {
	Value utils.Sequence[MeasresultssbIndexR15]
}
