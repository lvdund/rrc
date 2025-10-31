package ies

// HandoverPreparationInformation-v1320-IEs ::= SEQUENCE
type HandoverpreparationinformationV1320 struct {
	AsConfigV1320        *AsConfigV1320
	AsContextV1320       *AsContextV1320
	Noncriticalextension *HandoverpreparationinformationV1430
}
