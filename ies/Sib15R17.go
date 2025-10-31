package ies

import "rrc/utils"

// SIB15-r17 ::= SEQUENCE
// Extensible
type Sib15R17 struct {
	CommonplmnswithdisasterconditionR17 *[]PlmnIdentity              `lb:1,ub:maxPLMN`
	ApplicabledisasterinfolistR17       *[]ApplicabledisasterinfoR17 `lb:1,ub:maxPLMN`
	Latenoncriticalextension            *utils.OCTETSTRING
}
