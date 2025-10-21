package ies

import "rrc/utils"

// UE-EUTRA-CapabilityAddXDD-Mode-v1530 ::= SEQUENCE
type UeEutraCapabilityaddxddModeV1530 struct {
	NeighcellsiAcquisitionparametersV1530 *NeighcellsiAcquisitionparametersV1530
	ReducedcpLatencyR15                   *utils.ENUMERATED
}
