package ies

import "rrc/utils"

// MeasReportAppLayer-r17 ::= SEQUENCE
type MeasreportapplayerR17 struct {
	MeasconfigapplayeridR17        MeasconfigapplayeridR17
	MeasreportapplayercontainerR17 *utils.OCTETSTRING
	ApplayersessionstatusR17       *MeasreportapplayerR17ApplayersessionstatusR17
	RanVisiblemeasurementsR17      *RanVisiblemeasurementsR17
}
