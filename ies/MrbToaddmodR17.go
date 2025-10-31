package ies

import "rrc/utils"

// MRB-ToAddMod-r17 ::= SEQUENCE
// Extensible
type MrbToaddmodR17 struct {
	MbsSessionidR17    *TmgiR17
	MrbIdentityR17     MrbIdentityR17
	MrbIdentitynewR17  *MrbIdentityR17
	ReestablishpdcpR17 *utils.BOOLEAN
	RecoverpdcpR17     *utils.BOOLEAN
	PdcpConfigR17      *PdcpConfig
}
