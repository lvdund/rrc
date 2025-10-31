package ies

import "rrc/utils"

// SchedulingRequestResourceId ::= utils.INTEGER (1..maxNrofSR-Resources)
type Schedulingrequestresourceid struct {
	Value utils.INTEGER `lb:0,ub:maxNrofSRResources`
}
