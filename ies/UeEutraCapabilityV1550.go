package ies

// UE-EUTRA-Capability-v1550-IEs ::= SEQUENCE
type UeEutraCapabilityV1550 struct {
	NeighcellsiAcquisitionparametersV1550 *NeighcellsiAcquisitionparametersV1550
	PhylayerparametersV1550               PhylayerparametersV1550
	MacParametersV1550                    MacParametersV1550
	FddAddUeEutraCapabilitiesV1550        UeEutraCapabilityaddxddModeV1550
	TddAddUeEutraCapabilitiesV1550        UeEutraCapabilityaddxddModeV1550
	Noncriticalextension                  *UeEutraCapabilityV1560
}
