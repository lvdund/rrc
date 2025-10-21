package ies

import "rrc/utils"

// MeasResultListRSSI-SCG-r13 ::= SEQUENCE OF MeasResultRSSI-SCG-r13
// SIZE (1..maxServCell-r13)
type MeasresultlistrssiScgR13 struct {
	Value utils.Sequence[MeasresultrssiScgR13]
}
