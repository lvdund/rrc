package ies

// MeasResultIdleNR-r16 ::= SEQUENCE
// Extensible
type MeasresultidlenrR16 struct {
	CarrierfreqnrR16                ArfcnValuenrR15
	MeasresultspercelllistidlenrR16 []MeasresultspercellidlenrR16 `lb:1,ub:maxCellMeasIdleR15`
}
