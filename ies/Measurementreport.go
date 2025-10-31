package ies

import "rrc/utils"

// MeasurementReport-IEs ::= SEQUENCE
type Measurementreport struct {
	Measresults              Measresults
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *MeasurementreportIesNoncriticalextension
}
