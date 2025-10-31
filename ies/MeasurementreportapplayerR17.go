package ies

import "rrc/utils"

// MeasurementReportAppLayer-r17-IEs ::= SEQUENCE
type MeasurementreportapplayerR17 struct {
	MeasurementreportapplayerlistR17 MeasurementreportapplayerlistR17
	Latenoncriticalextension         *utils.OCTETSTRING
	Noncriticalextension             *MeasurementreportapplayerR17IesNoncriticalextension
}
