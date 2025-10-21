package ies

import "rrc/utils"

// RRCConnectionResume-v1510-IEs ::= SEQUENCE
type RrcconnectionresumeV1510Ies struct {
	SkCounterR15            *utils.INTEGER
	NrRadiobearerconfig1R15 *utils.OCTETSTRING
	NrRadiobearerconfig2R15 *utils.OCTETSTRING
	Noncriticalextension    *RrcconnectionresumeV1530Ies
}
