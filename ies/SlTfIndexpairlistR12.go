package ies

// SL-TF-IndexPairList-r12 ::= SEQUENCE OF SL-TF-IndexPair-r12
// SIZE (1..maxSL-TF-IndexPair-r12)
type SlTfIndexpairlistR12 struct {
	Value []SlTfIndexpairR12 `lb:1,ub:maxSLTfIndexpairR12`
}
