package ies

// MeasurementReport-r8-IEs ::= SEQUENCE
type MeasurementreportR8 struct {
	Measresults          Measresults
	Noncriticalextension *MeasurementreportV8a0
}
