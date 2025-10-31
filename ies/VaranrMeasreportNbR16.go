package ies

import "rrc/utils"

// VarANR-MeasReport-NB-r16 ::= SEQUENCE
type VaranrMeasreportNbR16 struct {
	PlmnIdentitylistR16   PlmnIdentitylist3R11
	ServcellidentityR16   Cellglobalideutra
	MeasresultservcellR16 MeasresultservcellNbR14
	RelativetimestampR16  utils.INTEGER        `lb:0,ub:95`
	MeasresultlistR16     []AnrMeasresultNbR16 `lb:1,ub:maxFreqANRNbR16`
}
