package ies

import "rrc/utils"

// MeasurementReport-r8-IEs ::= SEQUENCE
type MeasurementreportR8Ies struct {
	Measresults          Measresults
	Noncriticalextension *MeasurementreportV8a0Ies
}
