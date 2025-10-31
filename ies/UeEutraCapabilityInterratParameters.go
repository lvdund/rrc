package ies

// UE-EUTRA-Capability-interRAT-Parameters ::= SEQUENCE
type UeEutraCapabilityInterratParameters struct {
	Utrafdd       *IratParametersutraFdd
	Utratdd128    *IratParametersutraTdd128
	Utratdd384    *IratParametersutraTdd384
	Utratdd768    *IratParametersutraTdd768
	Geran         *IratParametersgeran
	Cdma2000Hrpd  *IratParameterscdma2000Hrpd
	Cdma20001xrtt *IratParameterscdma20001xrtt
}
