package ies

import "rrc/utils"

// UEAssistanceInformation-v1430-IEs-rlm-Report-r14-rlm-Event-r14 ::= ENUMERATED
type UeassistanceinformationV1430IesRlmReportR14RlmEventR14 struct {
	Value utils.ENUMERATED
}

const (
	UeassistanceinformationV1430IesRlmReportR14RlmEventR14EnumeratedNothing = iota
	UeassistanceinformationV1430IesRlmReportR14RlmEventR14EnumeratedEarlyoutofsync
	UeassistanceinformationV1430IesRlmReportR14RlmEventR14EnumeratedEarlyinsync
)
