package ies

import "rrc/utils"

// MobilityFromEUTRACommand-r8-IEs ::= SEQUENCE
type MobilityfromeutracommandR8 struct {
	CsFallbackindicator  utils.BOOLEAN
	Purpose              MobilityfromeutracommandR8IesPurpose
	Noncriticalextension *MobilityfromeutracommandV8a0
}
