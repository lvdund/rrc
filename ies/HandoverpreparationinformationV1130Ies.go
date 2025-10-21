package ies

import "rrc/utils"

// HandoverPreparationInformation-v1130-IEs ::= SEQUENCE
type HandoverpreparationinformationV1130Ies struct {
	AsContextV1130       *AsContextV1130
	Noncriticalextension *HandoverpreparationinformationV1250Ies
}
