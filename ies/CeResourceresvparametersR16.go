package ies

import "rrc/utils"

// CE-ResourceResvParameters-r16 ::= SEQUENCE
type CeResourceresvparametersR16 struct {
	SubframeresourceresvdlCeModeaR16   *utils.ENUMERATED
	SubframeresourceresvdlCeModebR16   *utils.ENUMERATED
	SubframeresourceresvulCeModeaR16   *utils.ENUMERATED
	SubframeresourceresvulCeModebR16   *utils.ENUMERATED
	SlotsymbolresourceresvdlCeModeaR16 *utils.ENUMERATED
	SlotsymbolresourceresvdlCeModebR16 *utils.ENUMERATED
	SlotsymbolresourceresvulCeModeaR16 *utils.ENUMERATED
	SlotsymbolresourceresvulCeModebR16 *utils.ENUMERATED
	SubcarrierpuncturingceModeaR16     *utils.ENUMERATED
	SubcarrierpuncturingceModebR16     *utils.ENUMERATED
}
