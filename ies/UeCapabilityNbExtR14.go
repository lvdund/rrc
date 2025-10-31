package ies

// UE-Capability-NB-Ext-r14-IEs ::= SEQUENCE
type UeCapabilityNbExtR14 struct {
	UeCategoryNbR14         *UeCapabilityNbExtR14IesUeCategoryNbR14
	MacParametersR14        *MacParametersNbR14
	PhylayerparametersV1430 *PhylayerparametersNbV1430
	RfParametersV1430       RfParametersNbV1430
	Noncriticalextension    *UeCapabilityNbV1440
}
