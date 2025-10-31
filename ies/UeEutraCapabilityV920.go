package ies

// UE-EUTRA-Capability-v920-IEs ::= SEQUENCE
type UeEutraCapabilityV920 struct {
	PhylayerparametersV920             PhylayerparametersV920
	InterratParametersgeranV920        IratParametersgeranV920
	InterratParametersutraV920         *IratParametersutraV920
	InterratParameterscdma2000V920     *IratParameterscdma20001xrttV920
	DevicetypeR9                       *UeEutraCapabilityV920IesDevicetypeR9
	CsgProximityindicationparametersR9 CsgProximityindicationparametersR9
	NeighcellsiAcquisitionparametersR9 NeighcellsiAcquisitionparametersR9
	SonParametersR9                    SonParametersR9
	Noncriticalextension               *UeEutraCapabilityV940
}
