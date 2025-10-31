package ies

// UE-NR-Capability-v1530 ::= SEQUENCE
type UeNrCapabilityV1530 struct {
	FddAddUeNrCapabilitiesV1530 *UeNrCapabilityaddxddModeV1530
	TddAddUeNrCapabilitiesV1530 *UeNrCapabilityaddxddModeV1530
	Dummy                       *UeNrCapabilityV1530Dummy
	InterratParameters          *InterratParameters
	Inactivestate               *UeNrCapabilityV1530Inactivestate
	Delaybudgetreporting        *UeNrCapabilityV1530Delaybudgetreporting
	Noncriticalextension        *UeNrCapabilityV1540
}
