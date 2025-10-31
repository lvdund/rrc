package ies

import "rrc/utils"

// RRCReconfigurationFailureSidelink-r16-IEs ::= SEQUENCE
type RrcreconfigurationfailuresidelinkR16 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcreconfigurationfailuresidelinkR16IesNoncriticalextension
}
