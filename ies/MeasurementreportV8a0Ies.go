package ies

import "rrc/utils"

// MeasurementReport-v8a0-IEs ::= SEQUENCE
type MeasurementreportV8a0Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
