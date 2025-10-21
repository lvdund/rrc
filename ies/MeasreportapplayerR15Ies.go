package ies

import "rrc/utils"

// MeasReportAppLayer-r15-IEs ::= SEQUENCE
type MeasreportapplayerR15Ies struct {
	MeasreportapplayercontainerR15 *utils.OCTETSTRING
	ServicetypeR15                 *utils.ENUMERATED
	Noncriticalextension           *MeasreportapplayerV1590Ies
}
