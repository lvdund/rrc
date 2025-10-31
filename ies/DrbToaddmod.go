package ies

import "rrc/utils"

// DRB-ToAddMod ::= SEQUENCE
// Extensible
type DrbToaddmod struct {
	Cnassociation   *DrbToaddmodCnassociation
	DrbIdentity     DrbIdentity
	Reestablishpdcp *utils.BOOLEAN
	Recoverpdcp     *utils.BOOLEAN
	PdcpConfig      *PdcpConfig
	DapsConfigR16   *utils.BOOLEAN `ext`
}
