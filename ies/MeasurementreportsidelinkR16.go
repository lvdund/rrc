package ies

import "rrc/utils"

// MeasurementReportSidelink-r16-IEs ::= SEQUENCE
type MeasurementreportsidelinkR16 struct {
	SlMeasresultsR16         SlMeasresultsR16
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *MeasurementreportsidelinkR16IesNoncriticalextension
}
