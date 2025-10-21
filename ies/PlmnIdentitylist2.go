package ies

import "rrc/utils"

// PLMN-IdentityList2 ::= SEQUENCE OF PLMN-Identity
// SIZE (1..5)
type PlmnIdentitylist2 struct {
	Value []PlmnIdentity
}
