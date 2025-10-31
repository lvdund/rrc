package ies

// SCGFailureInformation-IEs ::= SEQUENCE
type Scgfailureinformation struct {
	Failurereportscg     *Failurereportscg
	Noncriticalextension *ScgfailureinformationV1590
}
