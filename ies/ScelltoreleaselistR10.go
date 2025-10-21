package ies

import "rrc/utils"

// SCellToReleaseList-r10 ::= SEQUENCE OF SCellIndex-r10
// SIZE (1..maxSCell-r10)
type ScelltoreleaselistR10 struct {
	Value utils.Sequence[ScellindexR10]
}
