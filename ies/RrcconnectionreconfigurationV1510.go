package ies

import "rrc/utils"

// RRCConnectionReconfiguration-v1510-IEs ::= SEQUENCE
type RrcconnectionreconfigurationV1510 struct {
	NrConfigR15             *RrcconnectionreconfigurationV1510IesNrConfigR15
	SkCounterR15            *utils.INTEGER `lb:0,ub:65535`
	NrRadiobearerconfig1R15 *utils.OCTETSTRING
	NrRadiobearerconfig2R15 *utils.OCTETSTRING
	TdmPatternconfigR15     *TdmPatternconfigR15
	Noncriticalextension    *RrcconnectionreconfigurationV1530
}
