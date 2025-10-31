package ies

// SCGFailureInformation-r12-IEs ::= SEQUENCE
type ScgfailureinformationR12 struct {
	FailurereportscgR12  *FailurereportscgR12
	Noncriticalextension *ScgfailureinformationV12d0a
}
