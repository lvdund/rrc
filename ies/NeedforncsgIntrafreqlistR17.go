package ies

// NeedForNCSG-IntraFreqList-r17 ::= SEQUENCE OF NeedForNCSG-IntraFreq-r17
// SIZE (1.. maxNrofServingCells)
type NeedforncsgIntrafreqlistR17 struct {
	Value []NeedforncsgIntrafreqR17 `lb:1,ub:maxNrofServingCells`
}
