package ies

import "rrc/utils"

// SL-InterUE-CoordinationScheme2-r17-sl-SlotLevelResourceExclusion-r17 ::= ENUMERATED
type SlInterueCoordinationscheme2R17SlSlotlevelresourceexclusionR17 struct {
	Value utils.ENUMERATED
}

const (
	SlInterueCoordinationscheme2R17SlSlotlevelresourceexclusionR17EnumeratedNothing = iota
	SlInterueCoordinationscheme2R17SlSlotlevelresourceexclusionR17EnumeratedEnabled
)
