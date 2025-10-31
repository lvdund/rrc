package ies

// MBMSInterestIndication-v1310-IEs ::= SEQUENCE
type MbmsinterestindicationV1310 struct {
	MbmsServicesR13      *MbmsServicelistR13
	Noncriticalextension *MbmsinterestindicationV1540
}
