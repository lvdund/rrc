package ies

// LogMeasResultListBT-r15 ::= SEQUENCE OF LogMeasResultBT-r15
// SIZE (1..maxBT-IdReport-r15)
type LogmeasresultlistbtR15 struct {
	Value []LogmeasresultbtR15 `lb:1,ub:maxBTIdreportR15`
}
