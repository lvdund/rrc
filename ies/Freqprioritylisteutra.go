package ies

// FreqPriorityListEUTRA ::= SEQUENCE OF FreqPriorityEUTRA
// SIZE (1..maxFreq)
type Freqprioritylisteutra struct {
	Value []Freqpriorityeutra `lb:1,ub:maxFreq`
}
