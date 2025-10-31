package ies

// UEAssistanceInformation-v1610-IEs ::= SEQUENCE
type UeassistanceinformationV1610 struct {
	OverheatingassistanceV1610 *OverheatingassistanceV1610
	Noncriticalextension       *UeassistanceinformationV1610IesNoncriticalextension
}
