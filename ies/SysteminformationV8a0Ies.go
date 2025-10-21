package ies

import "rrc/utils"

// SystemInformation-v8a0-IEs ::= SEQUENCE
type SysteminformationV8a0Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
