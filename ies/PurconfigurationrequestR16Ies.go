package ies

import "rrc/utils"

// PURConfigurationRequest-r16-IEs ::= SEQUENCE
type PurconfigurationrequestR16Ies struct {
	PurConfigrequestR16      *interface{}
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
