package ies

import "rrc/utils"

// RRCConnectionSetupComplete-NB-v1430-IEs ::= SEQUENCE
type RrcconnectionsetupcompleteNbV1430Ies struct {
	GummeiTypeR14        *utils.ENUMERATED
	DcnIdR14             *utils.INTEGER
	Noncriticalextension *RrcconnectionsetupcompleteNbV1470Ies
}
