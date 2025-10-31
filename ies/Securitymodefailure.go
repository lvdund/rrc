package ies

import "rrc/utils"

// SecurityModeFailure-IEs ::= SEQUENCE
type Securitymodefailure struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *SecuritymodefailureIesNoncriticalextension
}
