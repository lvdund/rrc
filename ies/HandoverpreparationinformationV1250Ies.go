package ies

import "rrc/utils"

// HandoverPreparationInformation-v1250-IEs ::= SEQUENCE
type HandoverpreparationinformationV1250Ies struct {
	UeSupportedearfcnR12 *ArfcnValueeutraR9
	AsConfigV1250        *AsConfigV1250
	Noncriticalextension *HandoverpreparationinformationV1320Ies
}
