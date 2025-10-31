package ies

// SL-MeasIdToRemoveList-r16 ::= SEQUENCE OF SL-MeasId-r16
// SIZE (1..maxNrofSL-MeasId-r16)
type SlMeasidtoremovelistR16 struct {
	Value []SlMeasidR16 `lb:1,ub:maxNrofSLMeasidR16`
}
