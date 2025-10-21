package ies

import "rrc/utils"

// MobilityFromEUTRACommand-v930-IEs ::= SEQUENCE
type MobilityfromeutracommandV930Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *MobilityfromeutracommandV960Ies
}
