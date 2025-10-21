package ies

import "rrc/utils"

// MeasResultServFreqListNR-r15 ::= SEQUENCE OF MeasResultServFreqNR-r15
// SIZE (1..maxServCell-r13)
type MeasresultservfreqlistnrR15 struct {
	Value utils.Sequence[MeasresultservfreqnrR15]
}
