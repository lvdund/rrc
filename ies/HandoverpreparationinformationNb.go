package ies

// HandoverPreparationInformation-NB-IEs ::= SEQUENCE
type HandoverpreparationinformationNb struct {
	UeRadioaccesscapabilityinfoR13 UeCapabilityNbR13
	AsConfigR13                    AsConfigNb
	RrmConfigR13                   *RrmConfigNb
	AsContextR13                   *AsContextNb
	Noncriticalextension           *HandoverpreparationinformationNbV1380
}
