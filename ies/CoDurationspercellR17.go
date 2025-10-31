package ies

import "rrc/utils"

// CO-DurationsPerCell-r17 ::= SEQUENCE
type CoDurationspercellR17 struct {
	ServingcellidR17     Servcellindex
	PositionindciR17     utils.INTEGER `lb:0,ub:maxSFIDciPayloadsize1`
	SubcarrierspacingR17 Subcarrierspacing
	CoDurationlistR17    []CoDurationR17 `lb:1,ub:64`
}
