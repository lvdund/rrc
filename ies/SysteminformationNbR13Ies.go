package ies

import "rrc/utils"

// SystemInformation-NB-r13-IEs ::= SEQUENCE
// Extensible
type SysteminformationNbR13Ies struct {
	SibTypeandinfoR13        interface{}
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
