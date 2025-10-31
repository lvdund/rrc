package ies

// NR-PRS-MeasurementInfo-r16 ::= SEQUENCE
// Extensible
type NrPrsMeasurementinfoR16 struct {
	DlPrsPointaR16                  ArfcnValuenr
	NrMeasprsRepetitionandoffsetR16 NrPrsMeasurementinfoR16NrMeasprsRepetitionandoffsetR16
	NrMeasprsLengthR16              NrPrsMeasurementinfoR16NrMeasprsLengthR16
}
