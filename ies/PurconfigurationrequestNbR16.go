package ies

import "rrc/utils"

// PURConfigurationRequest-NB-r16-IEs ::= SEQUENCE
type PurconfigurationrequestNbR16 struct {
	PurConfigrequestR16      *PurConfigrequestNbR16
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *PurconfigurationrequestNbR16IesNoncriticalextension
}
