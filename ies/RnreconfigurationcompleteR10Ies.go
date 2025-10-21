package ies

import "rrc/utils"

// RNReconfigurationComplete-r10-IEs ::= SEQUENCE
type RnreconfigurationcompleteR10Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
