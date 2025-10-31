package ies

// UE-EUTRA-Capability-v1430-IEs ::= SEQUENCE
type UeEutraCapabilityV1430 struct {
	PhylayerparametersV1430            PhylayerparametersV1430
	UeCategorydlV1430                  *UeEutraCapabilityV1430IesUeCategorydlV1430
	UeCategoryulV1430                  *UeEutraCapabilityV1430IesUeCategoryulV1430
	UeCategoryulV1430b                 *UeEutraCapabilityV1430IesUeCategoryulV1430b
	MacParametersV1430                 *MacParametersV1430
	MeasparametersV1430                *MeasparametersV1430
	PdcpParametersV1430                *PdcpParametersV1430
	RlcParametersV1430                 RlcParametersV1430
	RfParametersV1430                  *RfParametersV1430
	LaaParametersV1430                 *LaaParametersV1430
	LwaParametersV1430                 *LwaParametersV1430
	LwipParametersV1430                *LwipParametersV1430
	OtherparametersV1430               OtherParametersV1430
	MmtelParametersR14                 *MmtelParametersR14
	MobilityparametersR14              *MobilityparametersR14
	CeParametersV1430                  CeParametersV1430
	FddAddUeEutraCapabilitiesV1430     *UeEutraCapabilityaddxddModeV1430
	TddAddUeEutraCapabilitiesV1430     *UeEutraCapabilityaddxddModeV1430
	MbmsParametersV1430                *MbmsParametersV1430
	SlParametersV1430                  *SlParametersV1430
	UeBasednetwperfmeasparametersV1430 *UeBasednetwperfmeasparametersV1430
	HighspeedenhparametersR14          *HighspeedenhparametersR14
	Noncriticalextension               *UeEutraCapabilityV1440
}
