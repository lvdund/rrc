package ies

import "rrc/utils"

// RRCEarlyDataComplete-v1590-IEs ::= SEQUENCE
type RrcearlydatacompleteV1590Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
