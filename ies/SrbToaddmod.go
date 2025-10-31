package ies

import "rrc/utils"

// SRB-ToAddMod ::= SEQUENCE
// Extensible
type SrbToaddmod struct {
	SrbIdentity      SrbIdentity
	Reestablishpdcp  *utils.BOOLEAN
	Discardonpdcp    *utils.BOOLEAN
	PdcpConfig       *PdcpConfig
	SrbIdentityV1700 *SrbIdentityV1700 `ext`
}
