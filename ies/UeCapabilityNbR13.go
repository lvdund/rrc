package ies

import "rrc/utils"

// UE-Capability-NB-r13 ::= SEQUENCE
type UeCapabilityNbR13 struct {
	AccessstratumreleaseR13 AccessstratumreleaseNbR13
	UeCategoryNbR13         *utils.ENUMERATED
	MultipledrbR13          *utils.ENUMERATED
	PdcpParametersR13       *PdcpParametersNbR13
	PhylayerparametersR13   PhylayerparametersNbR13
	RfParametersR13         RfParametersNbR13
	Dummy                   *interface{}
}
