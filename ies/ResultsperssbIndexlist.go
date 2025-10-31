package ies

// ResultsPerSSB-IndexList ::= SEQUENCE OF ResultsPerSSB-Index
// SIZE (1..maxNrofIndexesToReport2)
type ResultsperssbIndexlist struct {
	Value []ResultsperssbIndex `lb:1,ub:maxNrofIndexesToReport2`
}
