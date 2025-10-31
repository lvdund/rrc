package ies

import "rrc/utils"

// DRB-ToAddMod-NB-r13 ::= SEQUENCE
// Extensible
type DrbToaddmodNbR13 struct {
	EpsBeareridentityR13      *utils.INTEGER `lb:0,ub:15`
	DrbIdentityR13            DrbIdentity
	PdcpConfigR13             *PdcpConfigNbR13
	RlcConfigR13              *RlcConfigNbR13
	LogicalchannelidentityR13 *utils.INTEGER `lb:0,ub:10`
	LogicalchannelconfigR13   *LogicalchannelconfigNbR13
}
