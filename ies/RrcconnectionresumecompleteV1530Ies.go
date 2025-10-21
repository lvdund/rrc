package ies

import "rrc/utils"

// RRCConnectionResumeComplete-v1530-IEs ::= SEQUENCE
type RrcconnectionresumecompleteV1530Ies struct {
	LogmeasavailablebtR15      *utils.ENUMERATED
	LogmeasavailablewlanR15    *utils.ENUMERATED
	IdlemeasavailableR15       *utils.ENUMERATED
	FlightpathinfoavailableR15 *utils.ENUMERATED
	Noncriticalextension       *RrcconnectionresumecompleteV1610Ies
}
