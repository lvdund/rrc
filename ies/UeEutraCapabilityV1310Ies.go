package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1310-IEs ::= SEQUENCE
type UeEutraCapabilityV1310Ies struct {
	UeCategorydlV1310              *utils.ENUMERATED
	UeCategoryulV1310              *utils.ENUMERATED
	PdcpParametersV1310            PdcpParametersV1310
	RlcParametersV1310             RlcParametersV1310
	MacParametersV1310             *MacParametersV1310
	PhylayerparametersV1310        *PhylayerparametersV1310
	RfParametersV1310              *RfParametersV1310
	MeasparametersV1310            *MeasparametersV1310
	DcParametersV1310              *DcParametersV1310
	SlParametersV1310              *SlParametersV1310
	ScptmParametersR13             *ScptmParametersR13
	CeParametersR13                *CeParametersR13
	InterratParameterswlanR13      IratParameterswlanR13
	LaaParametersR13               *LaaParametersR13
	LwaParametersR13               *LwaParametersR13
	WlanIwParametersV1310          WlanIwParametersV1310
	LwipParametersR13              LwipParametersR13
	FddAddUeEutraCapabilitiesV1310 *UeEutraCapabilityaddxddModeV1310
	TddAddUeEutraCapabilitiesV1310 *UeEutraCapabilityaddxddModeV1310
	Noncriticalextension           *UeEutraCapabilityV1320Ies
}
