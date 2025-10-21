package ies

import "rrc/utils"

// RRCConnectionSetupComplete-v1250-IEs ::= SEQUENCE
type RrcconnectionsetupcompleteV1250Ies struct {
	MobilitystateR12         *utils.ENUMERATED
	MobilityhistoryavailR12  *utils.ENUMERATED
	LogmeasavailablembsfnR12 *utils.ENUMERATED
	Noncriticalextension     *RrcconnectionsetupcompleteV1320Ies
}
