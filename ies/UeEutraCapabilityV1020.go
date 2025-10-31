package ies

import "rrc/utils"

// UE-EUTRA-Capability-v1020-IEs ::= SEQUENCE
type UeEutraCapabilityV1020 struct {
	UeCategoryV1020                  *utils.INTEGER `lb:0,ub:8`
	PhylayerparametersV1020          *PhylayerparametersV1020
	RfParametersV1020                *RfParametersV1020
	MeasparametersV1020              *MeasparametersV1020
	Featuregroupindrel10R10          *utils.BITSTRING `lb:32,ub:32`
	InterratParameterscdma2000V1020  *IratParameterscdma20001xrttV1020
	UeBasednetwperfmeasparametersR10 *UeBasednetwperfmeasparametersR10
	InterratParametersutraTddV1020   *IratParametersutraTddV1020
	Noncriticalextension             *UeEutraCapabilityV1060
}
