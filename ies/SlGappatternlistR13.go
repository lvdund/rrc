package ies

// SL-GapPatternList-r13 ::= SEQUENCE OF SL-GapPattern-r13
// SIZE (1..maxSL-GP-r13)
type SlGappatternlistR13 struct {
	Value []SlGappatternR13 `lb:1,ub:maxSLGpR13`
}
