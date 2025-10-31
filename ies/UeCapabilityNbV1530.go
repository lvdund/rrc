package ies

// UE-Capability-NB-v1530-IEs ::= SEQUENCE
type UeCapabilityNbV1530 struct {
	EarlydataUpR15          *UeCapabilityNbV1530IesEarlydataUpR15
	RlcParametersR15        RlcParametersNbR15
	MacParametersV1530      MacParametersNbV1530
	PhylayerparametersV1530 *PhylayerparametersNbV1530
	TddUeCapabilityR15      *TddUeCapabilityNbR15
	Noncriticalextension    *UeCapabilityNbV15x0
}
