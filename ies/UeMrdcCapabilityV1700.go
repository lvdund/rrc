package ies

// UE-MRDC-Capability-v1700 ::= SEQUENCE
type UeMrdcCapabilityV1700 struct {
	MeasandmobparametersmrdcV1700 MeasandmobparametersmrdcV1700
	Noncriticalextension          *UeMrdcCapabilityV1730
}
