package ies

import "rrc/utils"

// RRCConnectionSetupComplete-v1430-IEs ::= SEQUENCE
type RrcconnectionsetupcompleteV1430Ies struct {
	DcnIdR14             *utils.INTEGER
	Noncriticalextension *RrcconnectionsetupcompleteV1530Ies
}
