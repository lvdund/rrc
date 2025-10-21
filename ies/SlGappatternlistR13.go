package ies

import "rrc/utils"

// SL-GapPatternList-r13 ::= SEQUENCE OF SL-GapPattern-r13
// SIZE (1..maxSL-GP-r13)
type SlGappatternlistR13 struct {
	Value []SlGappatternR13
}
