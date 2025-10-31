package ies

// UE-NR-Capability-v1540 ::= SEQUENCE
type UeNrCapabilityV1540 struct {
	SdapParameters              *SdapParameters
	Overheatingind              *UeNrCapabilityV1540Overheatingind
	ImsParameters               *ImsParameters
	Fr1AddUeNrCapabilitiesV1540 *UeNrCapabilityaddfrxModeV1540
	Fr2AddUeNrCapabilitiesV1540 *UeNrCapabilityaddfrxModeV1540
	Fr1Fr2AddUeNrCapabilities   *UeNrCapabilityaddfrxMode
	Noncriticalextension        *UeNrCapabilityV1550
}
