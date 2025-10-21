package ies

import "rrc/utils"

// MobilityFromEUTRACommand-v960-IEs ::= SEQUENCE
type MobilityfromeutracommandV960Ies struct {
	Bandindicator        *Bandindicatorgeran
	Noncriticalextension *MobilityfromeutracommandV1530Ies
}
