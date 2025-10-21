package ies

import "rrc/utils"

// SL-TF-IndexPairList-r12 ::= SEQUENCE OF SL-TF-IndexPair-r12
// SIZE (1..maxSL-TF-IndexPair-r12)
type SlTfIndexpairlistR12 struct {
	Value utils.Sequence[SlTfIndexpairR12]
}
