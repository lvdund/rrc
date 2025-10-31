package ies

// UE-Capability-NB-r13 ::= SEQUENCE
type UeCapabilityNbR13 struct {
	AccessstratumreleaseR13 AccessstratumreleaseNbR13
	UeCategoryNbR13         *UeCapabilityNbR13UeCategoryNbR13
	MultipledrbR13          *UeCapabilityNbR13MultipledrbR13
	PdcpParametersR13       *PdcpParametersNbR13
	PhylayerparametersR13   PhylayerparametersNbR13
	RfParametersR13         RfParametersNbR13
	Dummy                   *UeCapabilityNbR13Dummy
}
