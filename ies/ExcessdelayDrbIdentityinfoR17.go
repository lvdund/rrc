package ies

// ExcessDelay-DRB-IdentityInfo-r17 ::= SEQUENCE
type ExcessdelayDrbIdentityinfoR17 struct {
	DrbIdentitylist []DrbIdentity `lb:1,ub:maxDRB`
	Delaythreshold  ExcessdelayDrbIdentityinfoR17Delaythreshold
}
