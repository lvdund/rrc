package ies

import "rrc/utils"

// RRCReestablishmentComplete-IEs ::= SEQUENCE
type Rrcreestablishmentcomplete struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcreestablishmentcompleteV1610
}
