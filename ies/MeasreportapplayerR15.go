package ies

import "rrc/utils"

// MeasReportAppLayer-r15-IEs ::= SEQUENCE
type MeasreportapplayerR15 struct {
	MeasreportapplayercontainerR15 *utils.OCTETSTRING `lb:1,ub:8000`
	ServicetypeR15                 *MeasreportapplayerR15IesServicetypeR15
	Noncriticalextension           *MeasreportapplayerV1590
}
