package ies

import "rrc/utils"

// CO-DurationsPerCell-r16 ::= SEQUENCE
type CoDurationspercellR16 struct {
	ServingcellidR16     Servcellindex
	PositionindciR16     utils.INTEGER `lb:0,ub:maxSFIDciPayloadsize1`
	SubcarrierspacingR16 Subcarrierspacing
	CoDurationlistR16    []CoDurationR16 `lb:1,ub:64`
}
