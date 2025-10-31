package ies

import "rrc/utils"

// DRB-InfoSCG-r12 ::= SEQUENCE
// Extensible
type DrbInfoscgR12 struct {
	EpsBeareridentityR12 *utils.INTEGER `lb:0,ub:15`
	DrbIdentityR12       DrbIdentity
	DrbTypeR12           *DrbInfoscgR12DrbTypeR12
}
