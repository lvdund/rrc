package ies

import "rrc/utils"

// IKE-Identity-r13 ::= SEQUENCE
type IkeIdentityR13 struct {
	IdiR13 utils.OCTETSTRING
}
