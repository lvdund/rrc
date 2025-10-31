package ies

// NeedForGapsIntraFreqList-r16 ::= SEQUENCE OF NeedForGapsIntraFreq-r16
// SIZE (1.. maxNrofServingCells)
type NeedforgapsintrafreqlistR16 struct {
	Value []NeedforgapsintrafreqR16 `lb:1,ub:maxNrofServingCells`
}
