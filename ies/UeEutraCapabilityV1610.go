package ies

// UE-EUTRA-Capability-v1610-IEs ::= SEQUENCE
type UeEutraCapabilityV1610 struct {
	HighspeedenhparametersV1610           *HighspeedenhparametersV1610
	NeighcellsiAcquisitionparametersV1610 *NeighcellsiAcquisitionparametersV1610
	MbmsParametersV1610                   *MbmsParametersV1610
	PdcpParametersV1610                   *PdcpParametersV1610
	MacParametersV1610                    *MacParametersV1610
	PhylayerparametersV1610               *PhylayerparametersV1610
	MeasparametersV1610                   *MeasparametersV1610
	PurParametersR16                      *PurParametersR16
	Eutra5gcParametersV1610               *Eutra5gcParametersV1610
	OtherparametersV1610                  *OtherParametersV1610
	DlDedicatedmessagesegmentationR16     *UeEutraCapabilityV1610IesDlDedicatedmessagesegmentationR16
	MmtelParametersV1610                  MmtelParametersV1610
	IratParametersnrV1610                 *IratParametersnrV1610
	RfParametersV1610                     *RfParametersV1610
	MobilityparametersV1610               *MobilityparametersV1610
	UeBasednetwperfmeasparametersV1610    UeBasednetwperfmeasparametersV1610
	SlParametersV1610                     *SlParametersV1610
	FddAddUeEutraCapabilitiesV1610        *UeEutraCapabilityaddxddModeV1610
	TddAddUeEutraCapabilitiesV1610        *UeEutraCapabilityaddxddModeV1610
	Noncriticalextension                  *UeEutraCapabilityV1630
}
