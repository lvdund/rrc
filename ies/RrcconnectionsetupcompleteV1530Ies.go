package ies

import "rrc/utils"

// RRCConnectionSetupComplete-v1530-IEs ::= SEQUENCE
type RrcconnectionsetupcompleteV1530Ies struct {
	LogmeasavailablebtR15      *utils.ENUMERATED
	LogmeasavailablewlanR15    *utils.ENUMERATED
	IdlemeasavailableR15       *utils.ENUMERATED
	FlightpathinfoavailableR15 *utils.ENUMERATED
	Connectto5gcR15            *utils.ENUMERATED
	RegisteredamfR15           *RegisteredamfR15
	SNssaiListR15              *interface{}
	Ng5gSTmsiBitsR15           *interface{}
	Noncriticalextension       *RrcconnectionsetupcompleteV1540Ies
}
