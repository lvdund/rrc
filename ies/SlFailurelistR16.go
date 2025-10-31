package ies

// SL-FailureList-r16 ::= SEQUENCE OF SL-Failure-r16
// SIZE (1..maxNrofSL-Dest-r16)
type SlFailurelistR16 struct {
	Value []SlFailureR16 `lb:1,ub:maxNrofSLDestR16`
}
