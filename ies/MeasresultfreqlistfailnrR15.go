package ies

import "rrc/utils"

// MeasResultFreqListFailNR-r15 ::= SEQUENCE OF MeasResultFreqFailNR-r15
// SIZE (1..maxFreqNR-r15)
type MeasresultfreqlistfailnrR15 struct {
	Value []MeasresultfreqfailnrR15
}
