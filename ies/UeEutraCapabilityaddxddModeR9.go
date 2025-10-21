package ies

import "rrc/utils"

// UE-EUTRA-CapabilityAddXDD-Mode-r9 ::= SEQUENCE
// Extensible
type UeEutraCapabilityaddxddModeR9 struct {
	PhylayerparametersR9               *Phylayerparameters
	FeaturegroupindicatorsR9           *utils.BITSTRING
	Featuregroupindrel9addR9           *utils.BITSTRING
	InterratParametersgeranR9          *IratParametersgeran
	InterratParametersutraR9           *IratParametersutraV920
	InterratParameterscdma2000R9       *IratParameterscdma20001xrttV920
	NeighcellsiAcquisitionparametersR9 *NeighcellsiAcquisitionparametersR9
}
