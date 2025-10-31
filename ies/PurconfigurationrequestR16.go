package ies

import "rrc/utils"

// PURConfigurationRequest-r16-IEs ::= SEQUENCE
type PurconfigurationrequestR16 struct {
	PurConfigrequestR16      *PurconfigurationrequestR16IesPurConfigrequestR16
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *PurconfigurationrequestR16IesNoncriticalextension
}
