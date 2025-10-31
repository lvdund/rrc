package ies

import "rrc/utils"

// RateMatchPattern-patternType-bitmaps-symbolsInResourceBlock ::= CHOICE
const (
	RatematchpatternPatterntypeBitmapsSymbolsinresourceblockChoiceNothing = iota
	RatematchpatternPatterntypeBitmapsSymbolsinresourceblockChoiceOneslot
	RatematchpatternPatterntypeBitmapsSymbolsinresourceblockChoiceTwoslots
)

type RatematchpatternPatterntypeBitmapsSymbolsinresourceblock struct {
	Choice   uint64
	Oneslot  *utils.BITSTRING `lb:14,ub:14`
	Twoslots *utils.BITSTRING `lb:28,ub:28`
}
