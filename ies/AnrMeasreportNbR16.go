package ies

import "rrc/utils"

// ANR-MeasReport-NB-r16 ::= SEQUENCE
// Extensible
type AnrMeasreportNbR16 struct {
	ServcellidentityR16   *Cellglobalideutra
	MeasresultservcellR16 MeasresultservcellNbR14
	RelativetimestampR16  utils.INTEGER        `lb:0,ub:95`
	MeasresultlistR16     []AnrMeasresultNbR16 `lb:1,ub:maxFreqANRNbR16`
}
