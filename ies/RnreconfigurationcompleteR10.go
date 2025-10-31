package ies

import "rrc/utils"

// RNReconfigurationComplete-r10-IEs ::= SEQUENCE
type RnreconfigurationcompleteR10 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RnreconfigurationcompleteR10IesNoncriticalextension
}
