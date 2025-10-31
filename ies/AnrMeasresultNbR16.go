package ies

// ANR-MeasResult-NB-r16 ::= SEQUENCE
type AnrMeasresultNbR16 struct {
	CarrierfreqR16            CarrierfreqNbR13
	PhyscellidR16             *Physcellid
	MeasresultlastservcellR16 MeasresultservcellNbR14
	MeasresultR16             *NrsrpRangeNbR14
	CgiInfoR16                *AnrMeasresultNbR16CgiInfoR16
}
