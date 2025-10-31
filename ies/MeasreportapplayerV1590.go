package ies

import "rrc/utils"

// MeasReportAppLayer-v1590-IEs ::= SEQUENCE
type MeasreportapplayerV1590 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *MeasreportapplayerV1590IesNoncriticalextension
}
