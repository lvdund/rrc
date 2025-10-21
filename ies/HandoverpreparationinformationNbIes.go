package ies

import "rrc/utils"

// HandoverPreparationInformation-NB-IEs ::= SEQUENCE
type HandoverpreparationinformationNbIes struct {
	UeRadioaccesscapabilityinfoR13 UeCapabilityNbR13
	AsConfigR13                    AsConfigNb
	RrmConfigR13                   *RrmConfigNb
	AsContextR13                   *AsContextNb
	Noncriticalextension           *HandoverpreparationinformationNbV1380Ies
}
