package ies

import "rrc/utils"

// VarANR-MeasReport-NB-r16 ::= SEQUENCE
type VaranrMeasreportNbR16 struct {
	PlmnIdentitylistR16   PlmnIdentitylist3R11
	ServcellidentityR16   Cellglobalideutra
	MeasresultservcellR16 MeasresultservcellNbR14
	RelativetimestampR16  utils.INTEGER
	MeasresultlistR16     interface{}
}
