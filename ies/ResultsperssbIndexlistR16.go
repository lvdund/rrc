package ies

// ResultsPerSSB-IndexList-r16 ::= SEQUENCE OF ResultsPerSSB-IndexIdle-r16
// SIZE (1.. maxNrofIndexesToReport)
type ResultsperssbIndexlistR16 struct {
	Value []ResultsperssbIndexidleR16 `lb:1,ub:maxNrofIndexesToReport`
}
