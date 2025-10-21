package ies

import "rrc/utils"

// RRCConnectionReconfiguration-v1510-IEs ::= SEQUENCE
type RrcconnectionreconfigurationV1510Ies struct {
	NrConfigR15             *interface{}
	SkCounterR15            *utils.INTEGER
	NrRadiobearerconfig1R15 *utils.OCTETSTRING
	NrRadiobearerconfig2R15 *utils.OCTETSTRING
	TdmPatternconfigR15     *TdmPatternconfigR15
	Noncriticalextension    *RrcconnectionreconfigurationV1530Ies
}
