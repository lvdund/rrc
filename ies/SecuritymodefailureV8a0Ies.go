package ies

import "rrc/utils"

// SecurityModeFailure-v8a0-IEs ::= SEQUENCE
type SecuritymodefailureV8a0Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
