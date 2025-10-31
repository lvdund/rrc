package ies

// MBMSInterestIndication-v1540-IEs ::= SEQUENCE
type MbmsinterestindicationV1540 struct {
	MbmsRomInfolistR15   *[]MbmsRomInfoR15 `lb:1,ub:maxMBMSServicelistperueR13`
	Noncriticalextension *MbmsinterestindicationV1610
}
