package ies

import "rrc/utils"

// RateMatchPattern-patternType-bitmaps ::= SEQUENCE
// Extensible
type RatematchpatternPatterntypeBitmaps struct {
	Resourceblocks         utils.BITSTRING `lb:275,ub:275`
	Symbolsinresourceblock RatematchpatternPatterntypeBitmapsSymbolsinresourceblock
	Periodicityandpattern  *RatematchpatternPatterntypeBitmapsPeriodicityandpattern
}
