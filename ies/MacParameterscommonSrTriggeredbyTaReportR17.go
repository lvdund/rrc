package ies

import "rrc/utils"

// MAC-ParametersCommon-sr-TriggeredBy-TA-Report-r17 ::= ENUMERATED
type MacParameterscommonSrTriggeredbyTaReportR17 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonSrTriggeredbyTaReportR17EnumeratedNothing = iota
	MacParameterscommonSrTriggeredbyTaReportR17EnumeratedSupported
)
