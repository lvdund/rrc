package ies

// UE-NR-Capability-v1610 ::= SEQUENCE
type UeNrCapabilityV1610 struct {
	IndevicecoexindR16                *UeNrCapabilityV1610IndevicecoexindR16
	DlDedicatedmessagesegmentationR16 *UeNrCapabilityV1610DlDedicatedmessagesegmentationR16
	NrdcParametersV1610               *NrdcParametersV1610
	PowsavParametersR16               *PowsavParametersR16
	Fr1AddUeNrCapabilitiesV1610       *UeNrCapabilityaddfrxModeV1610
	Fr2AddUeNrCapabilitiesV1610       *UeNrCapabilityaddfrxModeV1610
	BhRlfIndicationR16                *UeNrCapabilityV1610BhRlfIndicationR16
	DirectsnAdditionfirstrrcIabR16    *UeNrCapabilityV1610DirectsnAdditionfirstrrcIabR16
	BapParametersR16                  *BapParametersR16
	ReferencetimeprovisionR16         *UeNrCapabilityV1610ReferencetimeprovisionR16
	SidelinkparametersR16             *SidelinkparametersR16
	HighspeedparametersR16            *HighspeedparametersR16
	MacParametersV1610                *MacParametersV1610
	McgrlfRecoveryviascgR16           *UeNrCapabilityV1610McgrlfRecoveryviascgR16
	ResumewithstoredmcgScellsR16      *UeNrCapabilityV1610ResumewithstoredmcgScellsR16
	ResumewithstoredscgR16            *UeNrCapabilityV1610ResumewithstoredscgR16
	ResumewithscgConfigR16            *UeNrCapabilityV1610ResumewithscgConfigR16
	UeBasedperfmeasParametersR16      *UeBasedperfmeasParametersR16
	SonParametersR16                  *SonParametersR16
	OndemandsibConnectedR16           *UeNrCapabilityV1610OndemandsibConnectedR16
	Noncriticalextension              *UeNrCapabilityV1640
}
