package ies

import "rrc/utils"

// DRB-ToAddMod ::= SEQUENCE
// Extensible
type DrbToaddmod struct {
	EpsBeareridentity      *utils.INTEGER `lb:0,ub:15`
	DrbIdentity            DrbIdentity
	PdcpConfig             *PdcpConfig
	RlcConfig              *RlcConfig
	Logicalchannelidentity *utils.INTEGER `lb:0,ub:10`
	Logicalchannelconfig   *Logicalchannelconfig
}
