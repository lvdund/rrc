package ies

import "rrc/utils"

// DedicatedSIBRequest-r16-IEs ::= SEQUENCE
type DedicatedsibrequestR16 struct {
	OndemandsibRequestlistR16 *DedicatedsibrequestR16IesOndemandsibRequestlistR16
	Latenoncriticalextension  *utils.OCTETSTRING
	Noncriticalextension      *DedicatedsibrequestR16IesNoncriticalextension
}
