package ies

// FailureInformation-r16-IEs ::= SEQUENCE
type FailureinformationR16 struct {
	FailedlogicalchannelidentityR16 *FailedlogicalchannelidentityR16
	FailuretypeR16                  *FailureinformationR16IesFailuretypeR16
	Noncriticalextension            *FailureinformationR16IesNoncriticalextension
}
