package ies

import "rrc/utils"

// SystemInformation-v8a0-IEs ::= SEQUENCE
type SysteminformationV8a0 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *SysteminformationV8a0IesNoncriticalextension
}
