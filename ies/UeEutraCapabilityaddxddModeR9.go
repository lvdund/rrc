package ies

import "rrc/utils"

// UE-EUTRA-CapabilityAddXDD-Mode-r9 ::= SEQUENCE
// Extensible
type UeEutraCapabilityaddxddModeR9 struct {
	PhylayerparametersR9               *Phylayerparameters
	FeaturegroupindicatorsR9           *utils.BITSTRING `lb:32,ub:32`
	Featuregroupindrel9addR9           *utils.BITSTRING `lb:32,ub:32`
	InterratParametersgeranR9          *IratParametersgeran
	InterratParametersutraR9           *IratParametersutraV920
	InterratParameterscdma2000R9       *IratParameterscdma20001xrttV920
	NeighcellsiAcquisitionparametersR9 *NeighcellsiAcquisitionparametersR9
}
