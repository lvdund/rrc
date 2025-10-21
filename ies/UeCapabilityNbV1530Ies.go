package ies

import "rrc/utils"

// UE-Capability-NB-v1530-IEs ::= SEQUENCE
type UeCapabilityNbV1530Ies struct {
	EarlydataUpR15          *utils.ENUMERATED
	RlcParametersR15        RlcParametersNbR15
	MacParametersV1530      MacParametersNbV1530
	PhylayerparametersV1530 *PhylayerparametersNbV1530
	TddUeCapabilityR15      *TddUeCapabilityNbR15
	Noncriticalextension    *UeCapabilityNbV15x0Ies
}
