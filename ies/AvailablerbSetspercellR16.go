package ies

import "rrc/utils"

// AvailableRB-SetsPerCell-r16 ::= SEQUENCE
type AvailablerbSetspercellR16 struct {
	ServingcellidR16 Servcellindex
	PositionindciR16 utils.INTEGER `lb:0,ub:maxSFIDciPayloadsize1`
}
