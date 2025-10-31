package ies

// UE-EUTRA-Capability-v1470-IEs ::= SEQUENCE
type UeEutraCapabilityV1470 struct {
	MbmsParametersV1470     *MbmsParametersV1470
	PhylayerparametersV1470 *PhylayerparametersV1470
	RfParametersV1470       *RfParametersV1470
	Noncriticalextension    *UeEutraCapabilityV14a0
}
