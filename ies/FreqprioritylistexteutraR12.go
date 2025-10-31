package ies

// FreqPriorityListExtEUTRA-r12 ::= SEQUENCE OF FreqPriorityEUTRA-r12
// SIZE (1..maxFreq)
type FreqprioritylistexteutraR12 struct {
	Value []FreqpriorityeutraR12 `lb:1,ub:maxFreq`
}
