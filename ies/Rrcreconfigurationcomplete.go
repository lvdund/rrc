package ies

import "rrc/utils"

// RRCReconfigurationComplete-IEs ::= SEQUENCE
type Rrcreconfigurationcomplete struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcreconfigurationcompleteV1530
}
