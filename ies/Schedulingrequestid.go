package ies

import "rrc/utils"

// SchedulingRequestId ::= utils.INTEGER (0..7)
type Schedulingrequestid struct {
	Value utils.INTEGER `lb:0,ub:7`
}
