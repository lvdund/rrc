package ies

import "rrc/utils"

// RRCEarlyDataRequest-v1590-IEs ::= SEQUENCE
type RrcearlydatarequestV1590 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcearlydatarequestV1610
}
