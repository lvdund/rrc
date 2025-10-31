package ies

// FailureInformation-v1610-IEs ::= SEQUENCE
type FailureinformationV1610 struct {
	FailureinfodapsR16   *FailureinfodapsR16
	Noncriticalextension *FailureinformationV1610IesNoncriticalextension
}
