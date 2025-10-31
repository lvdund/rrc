package ies

import "rrc/utils"

// MeasurementReport-v8a0-IEs ::= SEQUENCE
type MeasurementreportV8a0 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *MeasurementreportV8a0IesNoncriticalextension
}
