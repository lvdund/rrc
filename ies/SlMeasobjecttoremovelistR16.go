package ies

// SL-MeasObjectToRemoveList-r16 ::= SEQUENCE OF SL-MeasObjectId-r16
// SIZE (1..maxNrofSL-ObjectId-r16)
type SlMeasobjecttoremovelistR16 struct {
	Value []SlMeasobjectidR16 `lb:1,ub:maxNrofSLObjectidR16`
}
