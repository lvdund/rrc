package ies

import "rrc/utils"

// RRCConnectionReject-NB-r13-IEs ::= SEQUENCE
type RrcconnectionrejectNbR13 struct {
	ExtendedwaittimeR13      utils.INTEGER `lb:0,ub:1800`
	RrcSuspendindicationR13  *bool
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcconnectionrejectNbR13IesNoncriticalextension
}
