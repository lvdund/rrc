package ies

// PLMN-IdentityInfo-NB-r13 ::= SEQUENCE
type PlmnIdentityinfoNbR13 struct {
	PlmnIdentityR13                 PlmnIdentity
	CellreservedforoperatoruseR13   PlmnIdentityinfoNbR13CellreservedforoperatoruseR13
	AttachwithoutpdnConnectivityR13 *bool
}
