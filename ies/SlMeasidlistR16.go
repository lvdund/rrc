package ies

// SL-MeasIdList-r16 ::= SEQUENCE OF SL-MeasIdInfo-r16
// SIZE (1..maxNrofSL-MeasId-r16)
type SlMeasidlistR16 struct {
	Value []SlMeasidinfoR16 `lb:1,ub:maxNrofSLMeasidR16`
}
