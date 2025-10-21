package ies

import "rrc/utils"

// RRCEarlyDataComplete-NB-v1590-IEs ::= SEQUENCE
type RrcearlydatacompleteNbV1590Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
