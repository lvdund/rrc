package ies

import "rrc/utils"

// FailureInfoRLC-Bearer-failureType ::= ENUMERATED
type FailureinforlcBearerFailuretype struct {
	Value utils.ENUMERATED
}

const (
	FailureinforlcBearerFailuretypeEnumeratedNothing = iota
	FailureinforlcBearerFailuretypeEnumeratedRlc_Failure
	FailureinforlcBearerFailuretypeEnumeratedSpare3
	FailureinforlcBearerFailuretypeEnumeratedSpare2
	FailureinforlcBearerFailuretypeEnumeratedSpare1
)
