package ies

// HandoverPreparationInformation-IEs ::= SEQUENCE
type Handoverpreparationinformation struct {
	UeCapabilityratList  UeCapabilityratContainerlist
	Sourceconfig         *AsConfig
	RrmConfig            *RrmConfig
	AsContext            *AsContext
	Noncriticalextension *HandoverpreparationinformationIesNoncriticalextension
}
