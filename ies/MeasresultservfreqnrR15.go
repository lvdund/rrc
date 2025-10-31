package ies

// MeasResultServFreqNR-r15 ::= SEQUENCE
// Extensible
type MeasresultservfreqnrR15 struct {
	CarrierfreqR15             ArfcnValuenrR15
	MeasresultscellR15         *MeasresultcellnrR15
	MeasresultbestneighcellR15 *MeasresultcellnrR15
}
