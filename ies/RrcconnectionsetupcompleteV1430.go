package ies

import "rrc/utils"

// RRCConnectionSetupComplete-v1430-IEs ::= SEQUENCE
type RrcconnectionsetupcompleteV1430 struct {
	DcnIdR14             *utils.INTEGER `lb:0,ub:65535`
	Noncriticalextension *RrcconnectionsetupcompleteV1530
}
