package ies

// MeasResultFreqFailNR-r15 ::= SEQUENCE
// Extensible
type MeasresultfreqfailnrR15 struct {
	CarrierfreqR15        ArfcnValuenrR15
	MeasresultcelllistR15 *MeasresultcelllistnrR15
}
