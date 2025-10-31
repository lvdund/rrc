package ies

// FreqPriorityListDedicatedSlicing-r17 ::= SEQUENCE OF FreqPriorityDedicatedSlicing-r17
// SIZE (1.. maxFreq)
type FreqprioritylistdedicatedslicingR17 struct {
	Value []FreqprioritydedicatedslicingR17 `lb:1,ub:maxFreq`
}
