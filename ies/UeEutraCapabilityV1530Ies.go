package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1530-IEs ::= SEQUENCE
type UeEutraCapabilityV1530Ies struct {
	MeasparametersV1530                   *MeasparametersV1530
	OtherparametersV1530                  *OtherParametersV1530
	NeighcellsiAcquisitionparametersV1530 *NeighcellsiAcquisitionparametersV1530
	MacParametersV1530                    *MacParametersV1530
	PhylayerparametersV1530               *PhylayerparametersV1530
	RfParametersV1530                     *RfParametersV1530
	PdcpParametersV1530                   *PdcpParametersV1530
	UeCategorydlV1530                     *utils.INTEGER
	UeBasednetwperfmeasparametersV1530    *UeBasednetwperfmeasparametersV1530
	RlcParametersV1530                    *RlcParametersV1530
	SlParametersV1530                     *SlParametersV1530
	ExtendednumberofdrbsR15               *utils.ENUMERATED
	ReducedcpLatencyR15                   *utils.ENUMERATED
	LaaParametersV1530                    *LaaParametersV1530
	UeCategoryulV1530                     *utils.INTEGER
	FddAddUeEutraCapabilitiesV1530        *UeEutraCapabilityaddxddModeV1530
	TddAddUeEutraCapabilitiesV1530        *UeEutraCapabilityaddxddModeV1530
	Noncriticalextension                  *UeEutraCapabilityV1540Ies
}
