package ies

// NonContiguousUL-RA-WithinCC-List-r10 ::= SEQUENCE OF NonContiguousUL-RA-WithinCC-r10
// SIZE (1..maxBands)
type NoncontiguousulRaWithinccListR10 struct {
	Value []NoncontiguousulRaWithinccR10 `lb:1,ub:maxBands`
}
