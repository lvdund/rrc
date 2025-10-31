package ies

// MBS-FSAI-InterFreqList-r17 ::= SEQUENCE OF MBS-FSAI-InterFreq-r17
// SIZE (1..maxFreq)
type MbsFsaiInterfreqlistR17 struct {
	Value []MbsFsaiInterfreqR17 `lb:1,ub:maxFreq`
}
