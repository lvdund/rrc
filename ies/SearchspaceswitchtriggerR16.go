package ies

import "rrc/utils"

// SearchSpaceSwitchTrigger-r16 ::= SEQUENCE
type SearchspaceswitchtriggerR16 struct {
	ServingcellidR16 Servcellindex
	PositionindciR16 utils.INTEGER `lb:0,ub:maxSFIDciPayloadsize1`
}
