package ies

import "rrc/utils"

// SL-RemoteUE-RB-Identity-r17 ::= CHOICE
// Extensible
const (
	SlRemoteueRbIdentityR17ChoiceNothing = iota
	SlRemoteueRbIdentityR17ChoiceSrbIdentityR17
	SlRemoteueRbIdentityR17ChoiceDrbIdentityR17
)

type SlRemoteueRbIdentityR17 struct {
	Choice         uint64
	SrbIdentityR17 *utils.INTEGER `lb:0,ub:3`
	DrbIdentityR17 *DrbIdentity
}
