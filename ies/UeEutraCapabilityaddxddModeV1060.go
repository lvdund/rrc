package ies

import "rrc/utils"

// UE-EUTRA-CapabilityAddXDD-Mode-v1060 ::= SEQUENCE
// Extensible
type UeEutraCapabilityaddxddModeV1060 struct {
	PhylayerparametersV1060         *PhylayerparametersV1020
	Featuregroupindrel10V1060       *utils.BITSTRING
	InterratParameterscdma2000V1060 *IratParameterscdma20001xrttV1020
	InterratParametersutraTddV1060  *IratParametersutraTddV1020
}
