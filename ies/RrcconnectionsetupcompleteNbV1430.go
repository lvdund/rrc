package ies

import "rrc/utils"

// RRCConnectionSetupComplete-NB-v1430-IEs ::= SEQUENCE
type RrcconnectionsetupcompleteNbV1430 struct {
	GummeiTypeR14        *RrcconnectionsetupcompleteNbV1430IesGummeiTypeR14
	DcnIdR14             *utils.INTEGER `lb:0,ub:65535`
	Noncriticalextension *RrcconnectionsetupcompleteNbV1470
}
