package ies

import "rrc/utils"

// HandoverPreparationInformation-v1620-IEs ::= SEQUENCE
type HandoverpreparationinformationV1620Ies struct {
	AsContextV1620       *AsContextV1620
	Noncriticalextension *HandoverpreparationinformationV1630Ies
}
