package ies

import "rrc/utils"

// DRB-ToAddModSCG-r12 ::= SEQUENCE
// Extensible
type DrbToaddmodscgR12 struct {
	DrbIdentityR12               DrbIdentity
	DrbTypeR12                   *DrbToaddmodscgR12DrbTypeR12
	RlcConfigscgR12              *RlcConfig
	RlcConfigV1250               *RlcConfigV1250
	LogicalchannelidentityscgR12 *utils.INTEGER `lb:0,ub:10`
	LogicalchannelconfigscgR12   *Logicalchannelconfig
}
