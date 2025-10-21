package ies

import "rrc/utils"

// RRCConnectionReestablishmentComplete-v1530-IEs ::= SEQUENCE
type RrcconnectionreestablishmentcompleteV1530Ies struct {
	LogmeasavailablebtR15      *utils.ENUMERATED
	LogmeasavailablewlanR15    *utils.ENUMERATED
	FlightpathinfoavailableR15 *utils.ENUMERATED
	Noncriticalextension       *interface{}
}
