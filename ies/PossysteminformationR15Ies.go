package ies

import "rrc/utils"

// PosSystemInformation-r15-IEs ::= SEQUENCE
// Extensible
type PossysteminformationR15Ies struct {
	PossibTypeandinfoR15     interface{}
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
