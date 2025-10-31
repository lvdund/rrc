package ies

// LogMeasResultListBT-r16 ::= SEQUENCE OF LogMeasResultBT-r16
// SIZE (1..maxBT-IdReport-r16)
type LogmeasresultlistbtR16 struct {
	Value []LogmeasresultbtR16 `lb:1,ub:maxBTIdreportR16`
}
