package ies

import "rrc/utils"

// MeasResultListIdle-r15 ::= SEQUENCE OF MeasResultIdle-r15
// SIZE (1..maxIdleMeasCarriers-r15)
type MeasresultlistidleR15 struct {
	Value []MeasresultidleR15
}
