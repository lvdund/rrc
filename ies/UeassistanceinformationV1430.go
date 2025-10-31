package ies

// UEAssistanceInformation-v1430-IEs ::= SEQUENCE
type UeassistanceinformationV1430 struct {
	BwPreferenceR14             *BwPreferenceR14
	SpsAssistanceinformationR14 *UeassistanceinformationV1430IesSpsAssistanceinformationR14
	RlmReportR14                *UeassistanceinformationV1430IesRlmReportR14
	DelaybudgetreportR14        *DelaybudgetreportR14
	Noncriticalextension        *UeassistanceinformationV1450
}
