package ies

// SCGFailureInformationEUTRA-IEs ::= SEQUENCE
type Scgfailureinformationeutra struct {
	FailurereportscgEutra *FailurereportscgEutra
	Noncriticalextension  *ScgfailureinformationeutraV1590
}
