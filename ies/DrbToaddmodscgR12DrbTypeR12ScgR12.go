package ies

import "rrc/utils"

// DRB-ToAddModSCG-r12-drb-Type-r12-scg-r12 ::= SEQUENCE
type DrbToaddmodscgR12DrbTypeR12ScgR12 struct {
	EpsBeareridentityR12 *utils.INTEGER `lb:0,ub:15`
	PdcpConfigR12        *PdcpConfig
}
