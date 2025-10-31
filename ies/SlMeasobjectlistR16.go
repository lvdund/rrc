package ies

// SL-MeasObjectList-r16 ::= SEQUENCE OF SL-MeasObjectInfo-r16
// SIZE (1..maxNrofSL-ObjectId-r16)
type SlMeasobjectlistR16 struct {
	Value []SlMeasobjectinfoR16 `lb:1,ub:maxNrofSLObjectidR16`
}
