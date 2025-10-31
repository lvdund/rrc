package ies

// FreqPriorityListSlicing-r17 ::= SEQUENCE OF FreqPrioritySlicing-r17
// SIZE (1..maxFreqPlus1)
type FreqprioritylistslicingR17 struct {
	Value []FreqpriorityslicingR17 `lb:1,ub:maxFreqPlus1`
}
