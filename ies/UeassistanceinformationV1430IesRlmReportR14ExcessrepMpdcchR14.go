package ies

import "rrc/utils"

// UEAssistanceInformation-v1430-IEs-rlm-Report-r14-excessRep-MPDCCH-r14 ::= ENUMERATED
type UeassistanceinformationV1430IesRlmReportR14ExcessrepMpdcchR14 struct {
	Value utils.ENUMERATED
}

const (
	UeassistanceinformationV1430IesRlmReportR14ExcessrepMpdcchR14EnumeratedNothing = iota
	UeassistanceinformationV1430IesRlmReportR14ExcessrepMpdcchR14EnumeratedExcessrep1
	UeassistanceinformationV1430IesRlmReportR14ExcessrepMpdcchR14EnumeratedExcessrep2
)
