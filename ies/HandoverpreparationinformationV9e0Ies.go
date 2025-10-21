package ies

import "rrc/utils"

// HandoverPreparationInformation-v9e0-IEs ::= SEQUENCE
type HandoverpreparationinformationV9e0Ies struct {
	AsConfigV9e0         *AsConfigV9e0
	Noncriticalextension *HandoverpreparationinformationV1130Ies
}
