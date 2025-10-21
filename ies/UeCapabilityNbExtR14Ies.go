package ies

import "rrc/utils"

// UE-Capability-NB-Ext-r14-IEs ::= SEQUENCE
type UeCapabilityNbExtR14Ies struct {
	UeCategoryNbR14         *utils.ENUMERATED
	MacParametersR14        *MacParametersNbR14
	PhylayerparametersV1430 *PhylayerparametersNbV1430
	RfParametersV1430       RfParametersNbV1430
	Noncriticalextension    *UeCapabilityNbV1440Ies
}
