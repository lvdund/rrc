package ies

// UE-MRDC-Capability-v1610 ::= SEQUENCE
type UeMrdcCapabilityV1610 struct {
	MeasandmobparametersmrdcV1610 *MeasandmobparametersmrdcV1610
	GeneralparametersmrdcV1610    *GeneralparametersmrdcV1610
	PdcpParametersmrdcV1610       *PdcpParametersmrdcV1610
	Noncriticalextension          *UeMrdcCapabilityV1700
}
