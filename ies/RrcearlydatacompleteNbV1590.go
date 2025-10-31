package ies

import "rrc/utils"

// RRCEarlyDataComplete-NB-v1590-IEs ::= SEQUENCE
type RrcearlydatacompleteNbV1590 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcearlydatacompleteNbV1590IesNoncriticalextension
}
