package ies

import "rrc/utils"

// RRCEarlyDataRequest-NB-v1590-IEs ::= SEQUENCE
type RrcearlydatarequestNbV1590 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcearlydatarequestNbV1590IesNoncriticalextension
}
