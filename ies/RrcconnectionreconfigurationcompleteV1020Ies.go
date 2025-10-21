package ies

import "rrc/utils"

// RRCConnectionReconfigurationComplete-v1020-IEs ::= SEQUENCE
type RrcconnectionreconfigurationcompleteV1020Ies struct {
	RlfInfoavailableR10  *utils.ENUMERATED
	LogmeasavailableR10  *utils.ENUMERATED
	Noncriticalextension *RrcconnectionreconfigurationcompleteV1130Ies
}
