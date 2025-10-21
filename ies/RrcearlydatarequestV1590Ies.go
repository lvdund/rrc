package ies

import "rrc/utils"

// RRCEarlyDataRequest-v1590-IEs ::= SEQUENCE
type RrcearlydatarequestV1590Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcearlydatarequestV1610Ies
}
