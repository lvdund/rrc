package ies

import "rrc/utils"

// NonContiguousUL-RA-WithinCC-List-r10 ::= SEQUENCE OF NonContiguousUL-RA-WithinCC-r10
// SIZE (1..maxBands)
type NoncontiguousulRaWithinccListR10 struct {
	Value utils.Sequence[NoncontiguousulRaWithinccR10]
}
