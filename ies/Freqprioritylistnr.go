package ies

// FreqPriorityListNR ::= SEQUENCE OF FreqPriorityNR
// SIZE (1..maxFreq)
type Freqprioritylistnr struct {
	Value []Freqprioritynr `lb:1,ub:maxFreq`
}
