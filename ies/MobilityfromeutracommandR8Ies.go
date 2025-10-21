package ies

import "rrc/utils"

// MobilityFromEUTRACommand-r8-IEs ::= SEQUENCE
type MobilityfromeutracommandR8Ies struct {
	CsFallbackindicator  bool
	Purpose              interface{}
	Noncriticalextension *MobilityfromeutracommandV8a0Ies
}
