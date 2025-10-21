package ies

import "rrc/utils"

// FailureInformation-r16-IEs ::= SEQUENCE
type FailureinformationR16Ies struct {
	FailedlogicalchannelidentityR16 *FailedlogicalchannelidentityR16
	FailuretypeR16                  *utils.ENUMERATED
	Noncriticalextension            *interface{}
}
