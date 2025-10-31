package ies

import "rrc/utils"

// MobilityFromEUTRACommand-v8a0-IEs ::= SEQUENCE
type MobilityfromeutracommandV8a0 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *MobilityfromeutracommandV8d0
}
