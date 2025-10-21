package ies

import "rrc/utils"

// UEAssistanceInformation-v1430-IEs ::= SEQUENCE
type UeassistanceinformationV1430Ies struct {
	BwPreferenceR14             *BwPreferenceR14
	SpsAssistanceinformationR14 *interface{}
	RlmReportR14                *interface{}
	DelaybudgetreportR14        *DelaybudgetreportR14
	Noncriticalextension        *UeassistanceinformationV1450Ies
}
