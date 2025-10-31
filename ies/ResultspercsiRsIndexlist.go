package ies

// ResultsPerCSI-RS-IndexList ::= SEQUENCE OF ResultsPerCSI-RS-Index
// SIZE (1..maxNrofIndexesToReport2)
type ResultspercsiRsIndexlist struct {
	Value []ResultspercsiRsIndex `lb:1,ub:maxNrofIndexesToReport2`
}
