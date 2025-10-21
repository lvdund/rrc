package ies

import "rrc/utils"

// MeasResultListCBR-r14 ::= SEQUENCE OF MeasResultCBR-r14
// SIZE (1..maxCBR-Report-r14)
type MeasresultlistcbrR14 struct {
	Value []MeasresultcbrR14
}
