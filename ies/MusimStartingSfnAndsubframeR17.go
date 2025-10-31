package ies

import "rrc/utils"

// MUSIM-Starting-SFN-AndSubframe-r17 ::= SEQUENCE
type MusimStartingSfnAndsubframeR17 struct {
	StartingSfnR17      utils.INTEGER `lb:0,ub:1023`
	StartingsubframeR17 utils.INTEGER `lb:0,ub:9`
}
