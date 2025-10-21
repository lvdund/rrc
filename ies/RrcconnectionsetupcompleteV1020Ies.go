package ies

import "rrc/utils"

// RRCConnectionSetupComplete-v1020-IEs ::= SEQUENCE
type RrcconnectionsetupcompleteV1020Ies struct {
	GummeiTypeR10          *utils.ENUMERATED
	RlfInfoavailableR10    *utils.ENUMERATED
	LogmeasavailableR10    *utils.ENUMERATED
	RnSubframeconfigreqR10 *utils.ENUMERATED
	Noncriticalextension   *RrcconnectionsetupcompleteV1130Ies
}
