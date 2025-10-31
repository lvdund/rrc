package ies

import "rrc/utils"

// MobilityFromEUTRACommand-r9-IEs ::= SEQUENCE
// Extensible
type MobilityfromeutracommandR9 struct {
	CsFallbackindicator  utils.BOOLEAN
	Purpose              MobilityfromeutracommandR9IesPurpose
	Noncriticalextension *MobilityfromeutracommandV930
}
