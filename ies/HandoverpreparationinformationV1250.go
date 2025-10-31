package ies

// HandoverPreparationInformation-v1250-IEs ::= SEQUENCE
type HandoverpreparationinformationV1250 struct {
	UeSupportedearfcnR12 *ArfcnValueeutraR9
	AsConfigV1250        *AsConfigV1250
	Noncriticalextension *HandoverpreparationinformationV1320
}
