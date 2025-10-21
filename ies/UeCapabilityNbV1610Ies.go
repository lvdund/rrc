package ies

import "rrc/utils"

// UE-Capability-NB-v1610-IEs ::= SEQUENCE
type UeCapabilityNbV1610Ies struct {
	EarlysecurityreactivationR16 *utils.ENUMERATED
	EarlydataUp5gcR16            *utils.ENUMERATED
	PurParametersR16             *PurParametersNbR16
	MacParametersV1610           MacParametersNbV1610
	PhylayerparametersV1610      *PhylayerparametersNbV1610
	SonParametersR16             *SonParametersNbR16
	MeasparametersR16            MeasparametersNbR16
	TddUeCapabilityV1610         *TddUeCapabilityNbV1610
	Noncriticalextension         *interface{}
}
