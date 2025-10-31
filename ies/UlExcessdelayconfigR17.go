package ies

// UL-ExcessDelayConfig-r17 ::= SEQUENCE
type UlExcessdelayconfigR17 struct {
	ExcessdelayDrblistR17 []ExcessdelayDrbIdentityinfoR17 `lb:1,ub:maxDRB`
}
