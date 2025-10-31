package ies

// HandoverPreparationInformation-r8-IEs ::= SEQUENCE
type HandoverpreparationinformationR8 struct {
	UeRadioaccesscapabilityinfo UeCapabilityratContainerlist
	AsConfig                    *AsConfig
	RrmConfig                   *RrmConfig
	AsContext                   *AsContext
	Noncriticalextension        *HandoverpreparationinformationV920
}
