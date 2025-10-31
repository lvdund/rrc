package ies

import "rrc/utils"

// SRB-ToAddMod ::= SEQUENCE
// Extensible
type SrbToaddmod struct {
	SrbIdentity          utils.INTEGER `lb:0,ub:2`
	RlcConfig            *SrbToaddmodRlcConfig
	Logicalchannelconfig *SrbToaddmodLogicalchannelconfig
}
