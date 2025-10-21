package ies

import "rrc/utils"

// RegisteredAMF-r15 ::= SEQUENCE
type RegisteredamfR15 struct {
	PlmnIdentityR15  *PlmnIdentity
	AmfIdentifierR15 AmfIdentifierR15
}
