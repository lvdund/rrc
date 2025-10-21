package ies

import "rrc/utils"

// UE-EUTRA-Capability-v920-IEs ::= SEQUENCE
type UeEutraCapabilityV920Ies struct {
	PhylayerparametersV920             PhylayerparametersV920
	InterratParametersgeranV920        IratParametersgeranV920
	InterratParametersutraV920         *IratParametersutraV920
	InterratParameterscdma2000V920     *IratParameterscdma20001xrttV920
	DevicetypeR9                       *utils.ENUMERATED
	CsgProximityindicationparametersR9 CsgProximityindicationparametersR9
	NeighcellsiAcquisitionparametersR9 NeighcellsiAcquisitionparametersR9
	SonParametersR9                    SonParametersR9
	Noncriticalextension               *UeEutraCapabilityV940Ies
}
