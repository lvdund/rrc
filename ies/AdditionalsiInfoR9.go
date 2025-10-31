package ies

// AdditionalSI-Info-r9 ::= SEQUENCE
type AdditionalsiInfoR9 struct {
	CsgMemberstatusR9 *AdditionalsiInfoR9CsgMemberstatusR9
	CsgIdentityR9     *CsgIdentity
}
