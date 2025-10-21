package ies

import "rrc/utils"

// MeasResultListMBSFN-r12 ::= SEQUENCE OF MeasResultMBSFN-r12
// SIZE (1..maxMBSFN-Area)
type MeasresultlistmbsfnR12 struct {
	Value []MeasresultmbsfnR12
}
