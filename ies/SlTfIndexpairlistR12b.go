package ies

import "rrc/utils"

// SL-TF-IndexPairList-r12b ::= SEQUENCE OF SL-TF-IndexPair-r12b
// SIZE (1..maxSL-TF-IndexPair-r12)
type SlTfIndexpairlistR12b struct {
	Value []SlTfIndexpairR12b
}
