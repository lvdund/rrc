package ies

// PLMN-IdentityInfo-r15 ::= SEQUENCE
type PlmnIdentityinfoR15 struct {
	PlmnIdentity5gcR15               PlmnIdentityinfoR15PlmnIdentity5gcR15 `lb:1,ub:maxPLMNR11`
	CellreservedforoperatoruseR15    PlmnIdentityinfoR15CellreservedforoperatoruseR15
	CellreservedforoperatoruseCrsR15 PlmnIdentityinfoR15CellreservedforoperatoruseCrsR15
}
