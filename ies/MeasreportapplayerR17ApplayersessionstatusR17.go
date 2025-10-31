package ies

import "rrc/utils"

// MeasReportAppLayer-r17-appLayerSessionStatus-r17 ::= ENUMERATED
type MeasreportapplayerR17ApplayersessionstatusR17 struct {
	Value utils.ENUMERATED
}

const (
	MeasreportapplayerR17ApplayersessionstatusR17EnumeratedNothing = iota
	MeasreportapplayerR17ApplayersessionstatusR17EnumeratedStarted
	MeasreportapplayerR17ApplayersessionstatusR17EnumeratedStopped
)
