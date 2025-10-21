package ies

import "rrc/utils"

// TDD-UE-Capability-NB-v1610 ::= SEQUENCE
type TddUeCapabilityNbV1610 struct {
	SlotsymbolresourceresvdlR16 *utils.ENUMERATED
	SlotsymbolresourceresvulR16 *utils.ENUMERATED
	SubframeresourceresvdlR16   *utils.ENUMERATED
	SubframeresourceresvulR16   *utils.ENUMERATED
}
