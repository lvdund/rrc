package ies

// UECapabilityInformation-r8-IEs ::= SEQUENCE
type UecapabilityinformationR8 struct {
	UeCapabilityratContainerlist UeCapabilityratContainerlist
	Noncriticalextension         *UecapabilityinformationV8a0
}
