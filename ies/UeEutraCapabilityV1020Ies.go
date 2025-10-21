package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1020-IEs ::= SEQUENCE
type UeEutraCapabilityV1020Ies struct {
	UeCategoryV1020                  *utils.INTEGER
	PhylayerparametersV1020          *PhylayerparametersV1020
	RfParametersV1020                *RfParametersV1020
	MeasparametersV1020              *MeasparametersV1020
	Featuregroupindrel10R10          *utils.BITSTRING
	InterratParameterscdma2000V1020  *IratParameterscdma20001xrttV1020
	UeBasednetwperfmeasparametersR10 *UeBasednetwperfmeasparametersR10
	InterratParametersutraTddV1020   *IratParametersutraTddV1020
	Noncriticalextension             *UeEutraCapabilityV1060Ies
}
