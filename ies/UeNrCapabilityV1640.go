package ies

// UE-NR-Capability-v1640 ::= SEQUENCE
type UeNrCapabilityV1640 struct {
	RedirectatresumebynasR16               *UeNrCapabilityV1640RedirectatresumebynasR16
	PhyParameterssharedspectrumchaccessR16 *PhyParameterssharedspectrumchaccessR16
	Noncriticalextension                   *UeNrCapabilityV1650
}
