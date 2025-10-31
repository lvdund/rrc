package ies

// MeasResultNR-SL-r16 ::= SEQUENCE
// Extensible
type MeasresultnrSlR16 struct {
	MeasresultlistcbrNrR16 []MeasresultcbrNrR16 `lb:1,ub:maxNrofSLPooltomeasurenrR16`
}
