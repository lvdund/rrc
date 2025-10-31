package ies

import "rrc/utils"

// RateMatchPattern-patternType-bitmaps-periodicityAndPattern ::= CHOICE
const (
	RatematchpatternPatterntypeBitmapsPeriodicityandpatternChoiceNothing = iota
	RatematchpatternPatterntypeBitmapsPeriodicityandpatternChoiceN2
	RatematchpatternPatterntypeBitmapsPeriodicityandpatternChoiceN4
	RatematchpatternPatterntypeBitmapsPeriodicityandpatternChoiceN5
	RatematchpatternPatterntypeBitmapsPeriodicityandpatternChoiceN8
	RatematchpatternPatterntypeBitmapsPeriodicityandpatternChoiceN10
	RatematchpatternPatterntypeBitmapsPeriodicityandpatternChoiceN20
	RatematchpatternPatterntypeBitmapsPeriodicityandpatternChoiceN40
)

type RatematchpatternPatterntypeBitmapsPeriodicityandpattern struct {
	Choice uint64
	N2     *utils.BITSTRING `lb:2,ub:2`
	N4     *utils.BITSTRING `lb:4,ub:4`
	N5     *utils.BITSTRING `lb:5,ub:5`
	N8     *utils.BITSTRING `lb:8,ub:8`
	N10    *utils.BITSTRING `lb:10,ub:10`
	N20    *utils.BITSTRING `lb:20,ub:20`
	N40    *utils.BITSTRING `lb:40,ub:40`
}
