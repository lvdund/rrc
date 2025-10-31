package ies

// MBMSInterestIndication-v1610-IEs ::= SEQUENCE
type MbmsinterestindicationV1610 struct {
	MbmsRomInfolistR16   *[]MbmsRomInfoR16 `lb:1,ub:maxMBMSServicelistperueR13`
	Noncriticalextension *MbmsinterestindicationV1610IesNoncriticalextension
}
