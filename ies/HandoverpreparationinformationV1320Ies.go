package ies

import "rrc/utils"

// HandoverPreparationInformation-v1320-IEs ::= SEQUENCE
type HandoverpreparationinformationV1320Ies struct {
	AsConfigV1320        *AsConfigV1320
	AsContextV1320       *AsContextV1320
	Noncriticalextension *HandoverpreparationinformationV1430Ies
}
