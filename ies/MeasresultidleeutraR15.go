package ies

// MeasResultIdleEUTRA-r15 ::= SEQUENCE
// Extensible
type MeasresultidleeutraR15 struct {
	CarrierfreqR15 ArfcnValueeutraR9
	PhyscellidR15  Physcellid
	MeasresultR15  MeasresultidleeutraR15MeasresultR15
}
