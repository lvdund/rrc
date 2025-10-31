package ies

import "rrc/utils"

// RRCConnectionResume-v1510-IEs ::= SEQUENCE
type RrcconnectionresumeV1510 struct {
	SkCounterR15            *utils.INTEGER `lb:0,ub:65535`
	NrRadiobearerconfig1R15 *utils.OCTETSTRING
	NrRadiobearerconfig2R15 *utils.OCTETSTRING
	Noncriticalextension    *RrcconnectionresumeV1530
}
