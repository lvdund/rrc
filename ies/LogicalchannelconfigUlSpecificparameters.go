package ies

import "rrc/utils"

// LogicalChannelConfig-ul-SpecificParameters ::= SEQUENCE
type LogicalchannelconfigUlSpecificparameters struct {
	Priority            utils.INTEGER `lb:0,ub:16`
	Prioritisedbitrate  LogicalchannelconfigUlSpecificparametersPrioritisedbitrate
	Bucketsizeduration  LogicalchannelconfigUlSpecificparametersBucketsizeduration
	Logicalchannelgroup *utils.INTEGER `lb:0,ub:3`
}
