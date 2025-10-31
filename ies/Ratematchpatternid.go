package ies

import "rrc/utils"

// RateMatchPatternId ::= utils.INTEGER (0..maxNrofRateMatchPatterns-1)
type Ratematchpatternid struct {
	Value utils.INTEGER `lb:0,ub:maxNrofRateMatchPatterns1`
}
