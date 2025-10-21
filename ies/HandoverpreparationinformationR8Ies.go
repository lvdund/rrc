package ies

import "rrc/utils"

// HandoverPreparationInformation-r8-IEs ::= SEQUENCE
type HandoverpreparationinformationR8Ies struct {
	UeRadioaccesscapabilityinfo UeCapabilityratContainerlist
	AsConfig                    *AsConfig
	RrmConfig                   *RrmConfig
	AsContext                   *AsContext
	Noncriticalextension        *HandoverpreparationinformationV920Ies
}
