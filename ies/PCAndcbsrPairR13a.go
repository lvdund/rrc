package ies

import "rrc/utils"

// P-C-AndCBSR-Pair-r13a ::= SEQUENCE OF P-C-AndCBSR-r11
// SIZE (1..2)
type PCAndcbsrPairR13a struct {
	Value utils.Sequence[PCAndcbsrR11]
}
