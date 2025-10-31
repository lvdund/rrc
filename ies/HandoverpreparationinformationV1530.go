package ies

// HandoverPreparationInformation-v1530-IEs ::= SEQUENCE
type HandoverpreparationinformationV1530 struct {
	RanNotificationareainfoR15 *RanNotificationareainfoR15
	Noncriticalextension       *HandoverpreparationinformationV1540
}
