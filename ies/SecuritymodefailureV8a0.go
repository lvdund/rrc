package ies

import "rrc/utils"

// SecurityModeFailure-v8a0-IEs ::= SEQUENCE
type SecuritymodefailureV8a0 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *SecuritymodefailureV8a0IesNoncriticalextension
}
