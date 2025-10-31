package ies

// UEAssistanceInformation-v1530-IEs ::= SEQUENCE
type UeassistanceinformationV1530 struct {
	SpsAssistanceinformationV1530 *UeassistanceinformationV1530IesSpsAssistanceinformationV1530
	Noncriticalextension          *UeassistanceinformationV1610
}
