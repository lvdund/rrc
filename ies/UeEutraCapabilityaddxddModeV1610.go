package ies

import "rrc/utils"

// UE-EUTRA-CapabilityAddXDD-Mode-v1610 ::= SEQUENCE
type UeEutraCapabilityaddxddModeV1610 struct {
	PhylayerparametersV1610               *PhylayerparametersV1610
	PurParametersR16                      *PurParametersR16
	MeasparametersV1610                   *MeasparametersV1610
	Eutra5gcParametersV1610               *Eutra5gcParametersV1610
	IratParametersnrV1610                 *IratParametersnrV1610
	NeighcellsiAcquisitionparametersV1610 *NeighcellsiAcquisitionparametersV1610
	MobilityparametersV1610               *MobilityparametersV1610
}
