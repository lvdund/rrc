package ies

import "rrc/utils"

// MeasResultServFreqList-r10 ::= SEQUENCE OF MeasResultServFreq-r10
// SIZE (1..maxServCell-r10)
type MeasresultservfreqlistR10 struct {
	Value []MeasresultservfreqR10
}
