package ies

import "rrc/utils"

// MobilityFromEUTRACommand-v930-IEs ::= SEQUENCE
type MobilityfromeutracommandV930 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *MobilityfromeutracommandV960
}
