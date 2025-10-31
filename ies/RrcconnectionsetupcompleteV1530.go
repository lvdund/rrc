package ies

// RRCConnectionSetupComplete-v1530-IEs ::= SEQUENCE
type RrcconnectionsetupcompleteV1530 struct {
	LogmeasavailablebtR15      *bool
	LogmeasavailablewlanR15    *bool
	IdlemeasavailableR15       *bool
	FlightpathinfoavailableR15 *bool
	Connectto5gcR15            *bool
	RegisteredamfR15           *RegisteredamfR15
	SNssaiListR15              *[]SNssaiR15 `lb:1,ub:maxNrofSNssaiR15`
	Ng5gSTmsiBitsR15           *RrcconnectionsetupcompleteV1530IesNg5gSTmsiBitsR15
	Noncriticalextension       *RrcconnectionsetupcompleteV1540
}
