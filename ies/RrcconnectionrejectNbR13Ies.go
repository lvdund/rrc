package ies

import "rrc/utils"

// RRCConnectionReject-NB-r13-IEs ::= SEQUENCE
type RrcconnectionrejectNbR13Ies struct {
	ExtendedwaittimeR13      utils.INTEGER
	RrcSuspendindicationR13  *utils.ENUMERATED
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
