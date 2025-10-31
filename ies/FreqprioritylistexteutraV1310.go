package ies

// FreqPriorityListExtEUTRA-v1310 ::= SEQUENCE OF FreqPriorityEUTRA-v1310
// SIZE (1..maxFreq)
type FreqprioritylistexteutraV1310 struct {
	Value []FreqpriorityeutraV1310 `lb:1,ub:maxFreq`
}
