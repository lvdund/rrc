package ies

// FreqPriorityListNR-r15 ::= SEQUENCE OF FreqPriorityNR-r15
// SIZE (1..maxFreq)
type FreqprioritylistnrR15 struct {
	Value []FreqprioritynrR15 `lb:1,ub:maxFreq`
}
