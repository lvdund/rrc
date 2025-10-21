package ies

import "rrc/utils"

// RRCConnectionReconfigurationComplete-v1530-IEs ::= SEQUENCE
type RrcconnectionreconfigurationcompleteV1530Ies struct {
	LogmeasavailablebtR15      *utils.ENUMERATED
	LogmeasavailablewlanR15    *utils.ENUMERATED
	FlightpathinfoavailableR15 *utils.ENUMERATED
	Noncriticalextension       *interface{}
}
