package ies

import "rrc/utils"

// MobilityFromEUTRACommand-r9-IEs ::= SEQUENCE
// Extensible
type MobilityfromeutracommandR9Ies struct {
	CsFallbackindicator  bool
	Purpose              interface{}
	Noncriticalextension *MobilityfromeutracommandV930Ies
}
