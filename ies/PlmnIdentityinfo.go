package ies

// PLMN-IdentityInfo ::= SEQUENCE
type PlmnIdentityinfo struct {
	PlmnIdentity               PlmnIdentity
	Cellreservedforoperatoruse PlmnIdentityinfoCellreservedforoperatoruse
}
