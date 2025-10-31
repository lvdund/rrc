package ies

import "rrc/utils"

// RRCConnectionReconfiguration-v1130-IEs ::= SEQUENCE
type RrcconnectionreconfigurationV1130 struct {
	Systeminformationblocktype1dedicatedR11 *utils.OCTETSTRING
	Noncriticalextension                    *RrcconnectionreconfigurationV1250
}
