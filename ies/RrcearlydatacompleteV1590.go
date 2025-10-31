package ies

import "rrc/utils"

// RRCEarlyDataComplete-v1590-IEs ::= SEQUENCE
type RrcearlydatacompleteV1590 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcearlydatacompleteV1590IesNoncriticalextension
}
