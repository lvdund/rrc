package ies

// HandoverPreparationInformation-v1430-IEs ::= SEQUENCE
type HandoverpreparationinformationV1430 struct {
	AsConfigV1430         *AsConfigV1430
	MakebeforebreakreqR14 *bool
	Noncriticalextension  *HandoverpreparationinformationV1530
}
