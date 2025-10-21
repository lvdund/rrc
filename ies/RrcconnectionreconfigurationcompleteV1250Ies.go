package ies

import "rrc/utils"

// RRCConnectionReconfigurationComplete-v1250-IEs ::= SEQUENCE
type RrcconnectionreconfigurationcompleteV1250Ies struct {
	LogmeasavailablembsfnR12 *utils.ENUMERATED
	Noncriticalextension     *RrcconnectionreconfigurationcompleteV1430Ies
}
