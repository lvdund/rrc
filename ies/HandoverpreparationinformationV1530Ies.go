package ies

import "rrc/utils"

// HandoverPreparationInformation-v1530-IEs ::= SEQUENCE
type HandoverpreparationinformationV1530Ies struct {
	RanNotificationareainfoR15 *RanNotificationareainfoR15
	Noncriticalextension       *HandoverpreparationinformationV1540Ies
}
