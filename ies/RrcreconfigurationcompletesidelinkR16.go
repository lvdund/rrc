package ies

import "rrc/utils"

// RRCReconfigurationCompleteSidelink-r16-IEs ::= SEQUENCE
type RrcreconfigurationcompletesidelinkR16 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcreconfigurationcompletesidelinkV1710
}
