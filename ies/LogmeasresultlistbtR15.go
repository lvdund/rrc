package ies

import "rrc/utils"

// LogMeasResultListBT-r15 ::= SEQUENCE OF LogMeasResultBT-r15
// SIZE (1..maxBT-IdReport-r15)
type LogmeasresultlistbtR15 struct {
	Value utils.Sequence[LogmeasresultbtR15]
}
