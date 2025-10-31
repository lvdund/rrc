package ies

// UE-Capability-NB-v1610-IEs ::= SEQUENCE
type UeCapabilityNbV1610 struct {
	EarlysecurityreactivationR16 *UeCapabilityNbV1610IesEarlysecurityreactivationR16
	EarlydataUp5gcR16            *UeCapabilityNbV1610IesEarlydataUp5gcR16
	PurParametersR16             *PurParametersNbR16
	MacParametersV1610           MacParametersNbV1610
	PhylayerparametersV1610      *PhylayerparametersNbV1610
	SonParametersR16             *SonParametersNbR16
	MeasparametersR16            MeasparametersNbR16
	TddUeCapabilityV1610         *TddUeCapabilityNbV1610
	Noncriticalextension         *UeCapabilityNbV1610IesNoncriticalextension
}
